package new

import (
	"fmt"

	"github.com/gonrails/gonrails/cmds/helper"
)

func TouchSerializer(projectName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/serializers/book_serializer.go", helper.ProjectPath, projectName),
		fmt.Sprintf("%s/serializers/book_serializer.go", helper.TemplatePath),
		nil,
	)
}
