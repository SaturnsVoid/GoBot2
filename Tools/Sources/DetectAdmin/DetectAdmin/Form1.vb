Imports System.Security.Principal

Public Class Form1
    Private Sub Form1_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        MsgBox(Admin(), MsgBoxStyle.Information, "Is Admin?")
    End Sub
    Public Function Admin() As Boolean
        Try
            Dim Check As WindowsPrincipal = New WindowsPrincipal(WindowsIdentity.GetCurrent())
            Return Check.IsInRole(WindowsBuiltInRole.Administrator)
        Catch
            Return False
        End Try
    End Function
End Class
