package s3

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func filterEmptyString(str string) bool { return str != "" }

func stringToObjectIdentifier(str string) types.ObjectIdentifier {
	return types.ObjectIdentifier{Key: aws.String(str)}
}

func extractExtensionFromFilename(filename string) string {
	nameSplit := strings.Split(filename, ".")
	return nameSplit[len(nameSplit)-1]
}
