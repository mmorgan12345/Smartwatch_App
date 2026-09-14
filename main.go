package main

import (
	"context"
	"fmt"
)

// Main is the entrypoint for the DigitalOcean Go Function
func Main(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	// 1. Extract data sent by the smartwatch (e.g., heart rate)
	heartRate, ok := args["heart_rate"].(float64)
	if !ok {
		return map[string]interface{}{
			"body":       "Missing or invalid heart_rate parameter.",
			"statusCode": 400,
		}, nil
	}

	// 2. Process your smartwatch data (e.g., send to database, evaluate thresholds)
	message := fmt.Sprintf("Successfully logged heart rate: %.0f bpm", heartRate)

	// 3. Return the JSON response to the watch
	response := map[string]interface{}{
		"statusCode": 200,
		"headers": map[string]string{
			"Content-Type": "application/json",
		},
		"body": map[string]interface{}{
			"status":  "success",
			"message": message,
		},
	}

	return response, nil
}
