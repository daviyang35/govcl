//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package uac

// 这是一个空的文件，用于导入，也只用于测试用
// 生成的方式请前往wiki查看： https://gitee.com/ying32/govcl/wikis/pages?sort_id=410058&doc_id=102420

// uac的
// 使用方法为  import _ "github.com/ying32/govcl/pkgs/winappres/uac

// 只有windows下有效，里面包含默认的一些东西

// MANIFEST 文件
// 默认GoVCL的图标
// 默认空的版本信息

// 编译命令
// x86
// windres.exe -i app.rc -o defaultRes_windows_386.syso -F pe-i386
// x64
// windres.exe -i app.rc -o defaultRes_windows_amd64.syso -F pe-x86-64

// .manifest 文件
/*

<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0" xmlns:asmv3="urn:schemas-microsoft-com:asm.v3">
  <asmv3:application>
    <asmv3:windowsSettings xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">
      <!--<dpiAware>True/PM</dpiAware>-->
    </asmv3:windowsSettings>
  </asmv3:application>
  <dependency>
    <dependentAssembly>
      <assemblyIdentity
        type="win32"
        name="Microsoft.Windows.Common-Controls"
        version="6.0.0.0"
        publicKeyToken="6595b64144ccf1df"
        language="*"
        processorArchitecture="*"/>
    </dependentAssembly>
  </dependency>
  <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
      <requestedPrivileges>
        <requestedExecutionLevel
          level="asInvoker"
          uiAccess="false"
        />
        </requestedPrivileges>
    </security>
  </trustInfo>
<compatibility xmlns="urn:schemas-microsoft-com:compatibility.v1">
	<application>
		<!--The ID below indicates app support for Windows Vista -->
		<supportedOS Id="{e2011457-1546-43c5-a5fe-008deee3d3f0}"/>
		<!--The ID below indicates app support for Windows 7 -->
		<supportedOS Id="{35138b9a-5d96-4fbd-8e2d-a2440225f93a}"/>
		<!--The ID below indicates app support for Windows 8 -->
		<supportedOS Id="{4a2f28e3-53b9-4441-ba9c-d69d4a4a6e38}"/>
		<!--The ID below indicates app support for Windows 8.1 -->
		<supportedOS Id="{1f676c76-80e1-4239-95bb-83d0f6d0da78}"/>
		<!--The ID below indicates app support for Windows 10 -->
		<supportedOS Id="{8e0f7a12-bfb3-4fe8-b9a5-48fd50a15a9a}"/>
	</application>
</compatibility>
</assembly>

*/

// .rc文件
/*
1 MANIFEST "app.manifest"
MAINICON ICON "app.ico"

1 VERSIONINFO
FILEVERSION 1,0,0,0
PRODUCTVERSION 1,0,0,0
FILEOS 0x4
FILETYPE 0x1

BEGIN
    BLOCK "StringFileInfo"
    BEGIN
        BLOCK "040904B0"
        BEGIN
			VALUE "CompanyName", "<TODO:Company Name>"
			VALUE "FileDescription", ""
			VALUE "FileVersion", "1.0.0.0"
			VALUE "InternalName", ""
			VALUE "LegalCopyright", "<TODO:Copyright>"
			VALUE "OriginalFilename", ""
			VALUE "ProductName", ""
			VALUE "ProductVersion", "1.0.0.0"
        END
    END
    BLOCK "VarFileInfo"
    BEGIN
            VALUE "Translation", 0x0409, 0x04B0
    END
END

*/
