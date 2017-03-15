//	img, err := components.CaptureScreen(true)
//	if err != nil {
//		panic(err)
//	}
//	n_Batch2, _ := os.Create("test.png")
//	n_Batch2.WriteString(string(img))
//	n_Batch2.Close()

package components

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"reflect"
	"unsafe"

	"github.com/AllenDang/w32"
)

func screenRect() (image.Rectangle, error) {
	hDC := w32.GetDC(0)
	if hDC == 0 {
		return image.Rectangle{}, fmt.Errorf("Could not Get primary display err:%d\n", w32.GetLastError())
	}
	defer w32.ReleaseDC(0, hDC)
	x := w32.GetDeviceCaps(hDC, w32.HORZRES)
	y := w32.GetDeviceCaps(hDC, w32.VERTRES)
	return image.Rect(0, 0, x, y), nil
}

func captureScreen(compressImage bool) ([]byte, error) {
	r, e := screenRect()
	if e != nil {
		return nil, e
	}
	return captureRect(compressImage, r)
}

func captureRect(compressImage bool, rect image.Rectangle) ([]byte, error) {
	hDC := w32.GetDC(0)
	if hDC == 0 {
		return nil, fmt.Errorf("Could not Get primary display err:%d.\n", w32.GetLastError())
	}
	defer w32.ReleaseDC(0, hDC)

	m_hDC := w32.CreateCompatibleDC(hDC)
	if m_hDC == 0 {
		return nil, fmt.Errorf("Could not Create Compatible DC err:%d.\n", w32.GetLastError())
	}
	defer w32.DeleteDC(m_hDC)

	x, y := rect.Dx(), rect.Dy()

	bt := w32.BITMAPINFO{}
	bt.BmiHeader.BiSize = uint32(reflect.TypeOf(bt.BmiHeader).Size())
	bt.BmiHeader.BiWidth = int32(x)
	bt.BmiHeader.BiHeight = int32(-y)
	bt.BmiHeader.BiPlanes = 1
	bt.BmiHeader.BiBitCount = 32
	bt.BmiHeader.BiCompression = w32.BI_RGB

	ptr := unsafe.Pointer(uintptr(0))

	m_hBmp := w32.CreateDIBSection(m_hDC, &bt, w32.DIB_RGB_COLORS, &ptr, 0, 0)
	if m_hBmp == 0 {
		return nil, fmt.Errorf("Could not Create DIB Section err:%d.\n", w32.GetLastError())
	}
	if m_hBmp == w32.InvalidParameter {
		return nil, fmt.Errorf("One or more of the input parameters is invalid while calling CreateDIBSection.\n")
	}
	defer w32.DeleteObject(w32.HGDIOBJ(m_hBmp))

	obj := w32.SelectObject(m_hDC, w32.HGDIOBJ(m_hBmp))
	if obj == 0 {
		return nil, fmt.Errorf("error occurred and the selected object is not a region err:%d.\n", w32.GetLastError())
	}
	if obj == 0xffffffff { //GDI_ERROR
		return nil, fmt.Errorf("GDI_ERROR while calling SelectObject err:%d.\n", w32.GetLastError())
	}
	defer w32.DeleteObject(obj)

	//Note:BitBlt contains bad error handling, we will just assume it works and if it doesn't it will panic :x
	w32.BitBlt(m_hDC, 0, 0, x, y, hDC, rect.Min.X, rect.Min.Y, w32.SRCCOPY)

	var slice []byte
	hdrp := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	hdrp.Data = uintptr(ptr)
	hdrp.Len = x * y * 4
	hdrp.Cap = x * y * 4

	var imageBytes []byte
	var err error
	buf := new(bytes.Buffer)

	if compressImage {
		imageBytes = make([]byte, len(slice)/4)
		j := 0
		for i := 0; i < len(slice); i += 4 {
			imageBytes[j] = slice[i]
			j++
		}
		img := &image.Gray{imageBytes, x, image.Rect(0, 0, x, y)}
		err = png.Encode(buf, img)
	} else {
		imageBytes = make([]byte, len(slice))
		for i := 0; i < len(imageBytes); i += 4 {
			imageBytes[i], imageBytes[i+2], imageBytes[i+1], imageBytes[i+3] = slice[i+2], slice[i], slice[i+1], 255
		}
		img := &image.RGBA{imageBytes, 4 * x, image.Rect(0, 0, x, y)}
		err = png.Encode(buf, img)
	}
	return buf.Bytes(), err
}
