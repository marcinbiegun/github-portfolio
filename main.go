package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var config_yml = `
docs_subdir: '_docs'
raw_repo_url: 'https://raw.githubusercontent.com/marcinbiegun/creativecoding-sketches'
repo_url: 'https://github.com/marcinbiegun/creativecoding-sketches'

ignored_projects:
  - 'README.md'
  - '_docs'

root:
  text: |
    # creativecoding-sketches
    A collection of my creative coding doodles.

    ## Most interesting works

softwares:
  - name: 'Blender rrrrr'
    dir: 'Blender'
  - name: 'Cyrillo '
    dir: 'cyril'
    projects:
      - filename: ziom
        pinned: true
        text: loving this one
    project_dirs:
      - 'data'
    projects_in_dirs: true
`

// Todo add multiline variables

type ConfigSoftwareProject struct {
	Filename string
	Pinned   bool
	Text     string
}

type ConfigSoftware struct {
	Name             string
	Dir              string
	Projects         []ConfigSoftwareProject
	Project_dirs     []string
	Projects_in_dirs bool
}

type Config struct {
	Docs_subdir      string
	Raw_repo_url     string
	Repo_url         string
	Ignored_projects []string
	Root             struct {
		Text string
	}
	Softwares []ConfigSoftware `softwares`
}

func main() {
	config := Config{}
	err := yaml.Unmarshal([]byte(config_yml), &config)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("docs_subdir: %v\n", config.Docs_subdir)
	fmt.Printf("raw_repo_ur: %v\n", config.Raw_repo_url)
	fmt.Printf("repo_url: %v\n", config.Repo_url)

	for _, ignored_project := range config.Ignored_projects {
		fmt.Printf("  ignored_projects[]: %v\n", ignored_project)
	}

	fmt.Printf("root.text: %v\n", config.Root.Text)

	for _, software := range config.Softwares {
		fmt.Printf("  softwares[].name: %v\n", software.Name)
		fmt.Printf("  softwares[].dir: %v\n", software.Dir)
		fmt.Printf("  softwares[].project_dirs: %v\n", software.Project_dirs)
		fmt.Printf("  softwares[].projects_in_dirs: %v\n", software.Projects_in_dirs)

		for _, softwareProject := range software.Projects {
			fmt.Printf("  softwares[].projects[].filename: %v\n", softwareProject.Filename)
			fmt.Printf("  softwares[].projects[].pinned: %v\n", softwareProject.Pinned)
			fmt.Printf("  softwares[].projects[].text: %v\n", softwareProject.Text)
		}
	}
}
