package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type SortImageIds []*ecr.ImageIdentifier

func (c SortImageIds) Len() int {
	return len(c)
}

func (c SortImageIds) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortImageIds) Less(i, j int) bool {
	if c[i].ImageTag == nil {
		return true
	}
	if c[j].ImageTag == nil {
		return false
	}
	return strings.Compare(*c[i].ImageTag, *c[j].ImageTag) == -1
}

func GetSortedImageIds(region string, registryId string, repositoryName string) []*ecr.ImageIdentifier {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	ecrSvc := ecr.New(sess)

	done := false
	var imageIds []*ecr.ImageIdentifier
	params := &ecr.ListImagesInput{
		RepositoryName: aws.String(repositoryName),
		MaxResults:     aws.Int64(100),
		RegistryId:     aws.String(registryId),
	}
	for !done {
		resp, err := ecrSvc.ListImages(params)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		for _, imageID := range resp.ImageIds {
			imageIds = append(imageIds, imageID)
		}
		if resp.NextToken == nil {
			done = true
		} else {
			params.NextToken = resp.NextToken
		}
	}
	sort.Sort(SortImageIds(imageIds))

	return imageIds

}
