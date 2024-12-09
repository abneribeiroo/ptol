package cmd

import (
    "bytes"
    "fmt"
    "log"
    "os"
    "strings"
    "text/template"

    "github.com/spf13/cobra"
)

// Template for React component
const reactComponentTemplate = `
import React from 'react';

interface {{.ComponentName}}Props {
    // Add props here
}

const {{.ComponentName}}: React.FC<{{.ComponentName}}Props> = (props) => {
    return (
        <div>
            <h1>{{.ComponentName}} Component</h1>
            <p>This is a generated component.</p>
        </div>
    );
}

export default {{.ComponentName}};
`

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a React component",
    Run: func(cmd *cobra.Command, args []string) {
        var componentName string
        fmt.Print("Enter component name: ")
        fmt.Scanln(&componentName)

        if strings.TrimSpace(componentName) == "" {
            log.Fatal("Component name cannot be empty")
        }

        data := map[string]string{
            "ComponentName": componentName,
        }

        t, err := template.New("component").Parse(reactComponentTemplate)
        if err != nil {
            log.Fatal(err)
        }

        var buf bytes.Buffer
        err = t.Execute(&buf, data)
        if err != nil {
            log.Fatal(err)
        }

        fileName := fmt.Sprintf("%s.tsx", componentName)
        err = os.Mkdir("components", 0755)
        if err != nil && !os.IsExist(err) {
            log.Fatal(err)
        }

        err = os.WriteFile("components/"+fileName, buf.Bytes(), 0644)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Component %s generated successfully as %s\n", componentName, fileName)
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
}
