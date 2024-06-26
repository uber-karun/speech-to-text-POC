package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"google.golang.org/api/option"
)

func main() {
	// Path to your service account key file
	keyFile := "cred.json"

	// Path to the audio file you want to transcribe
	audioFile := "qwe.wav"

	ctx := context.Background()

	// Create a new Speech client
	client, err := speech.NewClient(ctx, option.WithCredentialsFile(keyFile))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Read the audio file into memory
	data, err := ioutil.ReadFile(audioFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Configure the request
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding: speechpb.RecognitionConfig_LINEAR16,
			// SampleRateHertz: 44100,
			LanguageCode: "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: data,
			},
		},
	}

	// Send the request to the API
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		log.Fatalf("Failed to recognize: %v", err)
	}

	fmt.Println(resp.TotalBilledTime.AsDuration(), len(resp.Results))
	// Print the results
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcription: %v\nConfidence: %v", alt.Transcript, alt.Confidence)
		}
	}
}
