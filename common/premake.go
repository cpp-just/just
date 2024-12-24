package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func CreateBaseFile(metaData MetaData) error {
	f, err := os.Create(".just/premake5.lua")
	if err != nil {
		return errors.New("failed to open the premake5.lua file")
	}

	defer f.Close()

	content := `require "cmake"` + "\n\n" +
		fmt.Sprintf(`workspace "%s"`, metaData.ProjectName) + "\n" +
		"\t" + `architecture "x64"` + "\n" +
		"\t" + `configurations { "Debug", "Release", "Dist" }` + "\n" +
		"\t" + fmt.Sprintf(`startproject "%s"`, metaData.ProjectName) + "\n\n" +
		`include "project.lua"`

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(content)
	if err != nil {
		return errors.New("failed to write to the premake5.lua file")
	}

	err = writer.Flush()
	if err != nil {
		return errors.New("failed to flush the buffer to the premake5.lua file")
	}

	return nil
}

func CreateProjectFile(metaData MetaData) error {
	f, err := os.Create(".just/project.lua")
	if err != nil {
		return errors.New("failed to open the premake5.lua file")
	}

	defer f.Close()

	content := fmt.Sprintf(`project "%s"`, metaData.ProjectName) + "\n" +
		"\t" + `kind "ConsoleApp"` + "\n" +
		"\t" + `language "C++"` + "\n" +
		"\t" + `cppdialect "C++17"` + "\n\n" +
		"\t" + `targetdir ("../bin/" .. outputdir .. "/%{prj.name}")` + "\n\n" +
		"\t" + `objdir ("../bin-int/" .. outputdir .. "/%{prj.name}")` + "\n" +
		"\t" + `files(get_source_files())` + "\n" +
		"\t" + "libdirs(get_lib_dirs())" + "\n" +
		"\t" + "links(get_links())" + "\n\n" +
		"\t" + `filter "configurations:Debug"` + "\n" +
		"\t\t" + `runtime "Debug"` + "\n" +
		"\t\t" + `symbols "On"` + "\n\n" +
		"\t" + `filter "configurations:Release"` + "\n" +
		"\t\t" + `runtime "Release"` + "\n" +
		"\t\t" + `optimize "On"` + "\n" +
		"\t\t" + `symbols "On"` + "\n\n" +
		"\t" + `filter "configurations:Dist"` + "\n" +
		"\t\t" + `runtime "Release"` + "\n" +
		"\t\t" + `optimize "On"` + "\n" +
		"\t\t" + `symbols "Off"` + "\n"

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(content)
	if err != nil {
		return errors.New("failed to write to the project.lua file")
	}

	err = writer.Flush()
	if err != nil {
		return errors.New("failed to flush the buffer to the project.lua file")
	}

	return nil
}
