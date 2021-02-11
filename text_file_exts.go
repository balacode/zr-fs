// -----------------------------------------------------------------------------
// ZR Library - File System Package                    zr-fs/[text_file_exts.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package fs

// TextFileExts is an array of text file extensions.
//
// *.asm *.asn *.asp *.aspx *.bas *.bat *.c *.cbp *.cc *.cfg *.classpath *.cls
// *.clw *.cpp *.cs *.csproj *.css *.ctl *.ctp *.cxx *.def *.dep *.dpr *.dsp
// *.dsw *.fd *.frm *.gitignore *.go *.h *.hc *.hh *.hhc *.hhk *.hhp *.hpp *.hs
// *.hta *.htm *.html *.hxx *.java *.jcl *.js *.json *.jsonp *.kt *.layout
// *.less *.log *.mak *.manifest *.meta *.odl *.pas *.pdm *.ph *.php *.pl *.plg
// *.pm *.prefs *.project *.properties *.py *.rb *.rc *.rc2 *.reg *.resx *.rgon
// *.rs *.rules *.sass *.scss *.settings *.sh *.shtml *.sln *.sql *.svg *.tlh
// *.tli *.ts *.txt *.user *.vbp *.vbw *.vcp *.vcproj *.vcw *.vcxproj
// *.workspace *.xaml *.xhtml *.xml *.xs *.xsd *.xsx *.yml
//
var TextFileExts = []string{ // TODO: use upper case
	//
	// Web-related Files
	"asp",
	"aspx",
	"css",
	"ctp",
	"hta",
	"htm",
	"html",
	"js",
	"less",
	"php",
	"sass",
	"scss",
	"shtml",
	"xhtml",
	//
	// Android/Java/Kotlin Files
	"classpath",
	"java",
	"kt",
	"prefs",
	"project",
	"properties",
	//
	// C/C++ Files
	"c",
	"cc",
	"cpp",
	"cxx",
	"fd",
	"h",
	"hc",
	"hh",
	"hpp",
	"hxx",
	"odl",
	"plg",
	"workspace",
	//
	// C# Files
	"cs",
	"user",
	"xsd",
	"xsx",
	// GIT
	"gitignore",
	//
	// Visual Basic
	"bas",
	"cls",
	"ctl",
	"dep",
	"frm",
	"pdm",
	//
	// Visual Studio / Other Project Files
	"cbp",
	"clw",
	"csproj",
	"def",
	"dsp",
	"dsw",
	"layout",
	"manifest",
	"rc",
	"rc2",
	"resx",
	"sln",
	"vbp",
	"vbw",
	"vcp", // Microsoft embedded Visual Tools Project
	"vcproj",
	"vcw", // Microsoft embedded Visual Tools Workspace File
	"vcxproj",
	//
	// Help Source
	"hhc",
	"hhk",
	"hhp",
	//
	// Markup
	"json",
	"jsonp",
	"rgon",
	"svg",
	"xaml",
	"xml",
	"yml",
	//
	// Miscellaneous
	"asm",
	"asn",
	"bat", // Windows batch file
	"cfg", // configuration file
	"dpr", // Delphi project file
	"go",  // Go Language source file
	"hs",  // Haskell source file
	"jcl", // job class (rarely-used file type)
	"log",
	"mak", // MAKE file
	"meta",
	"pas", // Pascal source file
	"ph",  // Perl source file
	"pl",  // Perl source file
	"pm",  // Perl source file
	"py",  // Python source file
	"rb",  // Ruby source file
	"reg",
	"rs", // Rust source file
	"rules",
	"settings",
	"sh",
	"sql",
	"tlh", // MSVC type library header (generated)
	"tli", // MSVC type library wrapper implementations (generated)
	"ts",  // TypeScript files
	"txt",
	"xs", // Perl extensions interface file
} //                                                                TextFileExts

//end
