package huggingface

import (
	"errors"

	"github.com/theotw/polyglot-llm/pkg/model"
	"github.com/theotw/polyglot-llm/pkg/utils"
)

const unsupportedAudioMessage = "huggingface provider does not currently support audio transcription in this library; use the openai or gemini provider"

func NewAudioTranscriptionGenerator(filePath string, opts model.AudioOptions) (model.AudioTranscriptionGenerator, error) {
	_ = filePath
	_ = opts
	return nil, utils.WrapIfNotNil(errors.New(unsupportedAudioMessage))
}
