package main

import (
	"fmt"
	bb "go_tutorial/bit_bool"
)

func testingBitBoolImplementation() {
	fmt.Println("3. Bit Boolean Pattern for 0b00001001")
	var enabledFeatures byte = 0b00000000
	enabledFeatures = bb.AddFeature(enabledFeatures, bb.FeatureAFlag) // enable A
	enabledFeatures = bb.AddFeature(enabledFeatures, bb.FeatureDFlag) // enable D

	fmt.Printf("Checking if Features A enabled: %v\n", bb.CheckIfFeatureEnabled(enabledFeatures, bb.FeatureAFlag)) // should be true
	fmt.Printf("Checking if Features B enabled: %v\n", bb.CheckIfFeatureEnabled(enabledFeatures, bb.FeatureBFlag)) // should be false
	fmt.Printf("Checking if Features D enabled: %v\n", bb.CheckIfFeatureEnabled(enabledFeatures, bb.FeatureDFlag)) // should be true
}

func main() {
	testingBitBoolImplementation()
}
