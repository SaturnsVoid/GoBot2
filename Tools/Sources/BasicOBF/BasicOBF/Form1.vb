Imports System.IO
Imports System.Security.Cryptography
Imports System.Text
Imports System.Text.RegularExpressions

Public Class Form1
    Dim tmpString As String
    Private Sub Button4_Click(sender As Object, e As EventArgs) Handles Button4.Click
        Try
            RichTextBox4.Clear()
            Dim col As MatchCollection = Regex.Matches(RichTextBox3.Text, """(.*?)""")
            For Each m As Match In col
                Dim g As Group = m.Groups(1)
                If CheckBox2.Checked = True Then
                    tmpString = ""
                    For i = 0 To g.Value.Length - 1
                        tmpString &= CChar(Char.ConvertFromUtf32(Char.ConvertToUtf32(g.Value(i), 0) + 1))
                    Next
                    RichTextBox4.Text &= "deobfuscate(""" + tmpString + """), " & vbNewLine
                Else
                    tmpString = ""
                    For i = 0 To g.Value.Length - 1
                        tmpString &= CChar(Char.ConvertFromUtf32(Char.ConvertToUtf32(g.Value(i), 0) + 1))
                    Next
                    RichTextBox4.Text &= """" + tmpString + """, " & vbNewLine
                End If
            Next
        Catch ex As Exception

        End Try
    End Sub

    Private Sub Button3_Click(sender As Object, e As EventArgs) Handles Button3.Click
        Try
            RichTextBox4.Clear()
            Dim col As MatchCollection = Regex.Matches(RichTextBox3.Text, """(.*?)""")
            For Each m As Match In col
                Dim g As Group = m.Groups(1)
                tmpString = ""
                For i = 0 To g.Value.Length - 1
                    tmpString &= CChar(Char.ConvertFromUtf32(Char.ConvertToUtf32(g.Value(i), 0) - 1))
                Next
                RichTextBox4.Text &= """" + tmpString + """, " & vbNewLine
            Next
        Catch ex As Exception

        End Try
    End Sub

    Private Sub Button2_Click_1(sender As Object, e As EventArgs) Handles Button2.Click
        Try
            RichTextBox5.Text = ""
            tmpString = ""
            For i = 0 To RichTextBox6.Text.Length - 1
                tmpString &= CChar(Char.ConvertFromUtf32(Char.ConvertToUtf32(RichTextBox6.Text(i), 0) - 1))
            Next

            RichTextBox5.Text &= tmpString
        Catch ex As Exception

        End Try
    End Sub

    Private Sub Button1_Click_1(sender As Object, e As EventArgs) Handles Button1.Click
        Try
            RichTextBox5.Text = ""
            tmpString = ""
            For i = 0 To RichTextBox6.Text.Length - 1
                tmpString &= CChar(Char.ConvertFromUtf32(Char.ConvertToUtf32(RichTextBox6.Text(i), 0) + 1))
            Next
            If CheckBox1.Checked = True Then
                RichTextBox5.Text &= "deobfuscate(""" & tmpString & """)"
            Else
                RichTextBox5.Text &= tmpString
            End If
        Catch ex As Exception

        End Try
    End Sub

    Private Sub Form1_Load(sender As Object, e As EventArgs) Handles MyBase.Load

    End Sub

    Private Sub Button8_Click(sender As Object, e As EventArgs) Handles Button8.Click
        If OpenFileDialog1.ShowDialog() = System.Windows.Forms.DialogResult.OK Then
            TextBox11.Text = OpenFileDialog1.FileName
        End If
    End Sub

    Private Sub Button9_Click(sender As Object, e As EventArgs) Handles Button9.Click
        If RadioButton1.Checked = True Then 'file
            Try
                RichTextBox2.Text = Convert.ToBase64String(System.IO.File.ReadAllBytes(TextBox11.Text))
            Catch ex As Exception

            End Try
        Else
            Try
                Dim byt As Byte() = System.Text.Encoding.UTF8.GetBytes(RichTextBox1.Text)
                RichTextBox2.Text = Convert.ToBase64String(byt)
            Catch ex As Exception

            End Try
        End If
    End Sub

    Private Sub Button10_Click(sender As Object, e As EventArgs) Handles Button10.Click
        If RadioButton1.Checked = True Then 'file
            If SaveFileDialog1.ShowDialog() = System.Windows.Forms.DialogResult.OK Then
                Try
                    Dim binaryData() As Byte = Convert.FromBase64String(RichTextBox1.Text)
                    Dim fs As New FileStream(SaveFileDialog1.FileName, FileMode.CreateNew)
                    fs.Write(binaryData, 0, binaryData.Length)
                    fs.Close()
                Catch ex As Exception

                End Try
            End If
        Else
            Try
                Dim b As Byte() = Convert.FromBase64String(RichTextBox1.Text)
                RichTextBox2.Text = System.Text.Encoding.UTF8.GetString(b)
            Catch ex As Exception

            End Try
        End If
    End Sub

    Private Sub Button5_Click(sender As Object, e As EventArgs) Handles Button5.Click
        TextBox1.Text = System.Guid.NewGuid.ToString
    End Sub

    Private Sub Button11_Click(sender As Object, e As EventArgs) Handles Button11.Click
        If OpenFileDialog1.ShowDialog() = System.Windows.Forms.DialogResult.OK Then
            TextBox2.Text = OpenFileDialog1.FileName
        End If
    End Sub

    Private Sub Button7_Click(sender As Object, e As EventArgs) Handles Button7.Click
        If RadioButton4.Checked = True Then 'file
            Try
                TextBox3.Text = GetFileHash(TextBox2.Text).ToLower
            Catch ex As Exception

            End Try
        Else
            Try
                TextBox3.Text = GetHash(RichTextBox8.Text).ToLower
            Catch ex As Exception

            End Try
        End If
    End Sub
    Function GetHash(theInput As String) As String
        Using hasher As MD5 = MD5.Create()
            Dim dbytes As Byte() =
                 hasher.ComputeHash(Encoding.UTF8.GetBytes(theInput))
            Dim sBuilder As New StringBuilder()
            For n As Integer = 0 To dbytes.Length - 1
                sBuilder.Append(dbytes(n).ToString("X2"))
            Next n
            Return sBuilder.ToString()
        End Using

    End Function
    Function GetFileHash(ByVal filename As String) As String
        Dim md5 As New MD5CryptoServiceProvider
        Dim f As New FileStream(filename, FileMode.Open, FileAccess.Read, FileShare.Read, &H2000)
        md5.ComputeHash(f)
        Dim hash As Byte() = md5.Hash
        Dim buff As New StringBuilder
        For Each hashByte As Byte In hash
            buff.Append(String.Format("{0:X2}", hashByte))
        Next
        f.Close()
        Return buff.ToString()
    End Function
End Class
