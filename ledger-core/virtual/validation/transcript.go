package validation

import (
	"github.com/insolar/assured-ledger/ledger-core/rms"
)

type Transcript struct {
	Entries []TranscriptEntry
}

func NewTranscript() Transcript {
	return Transcript{
		Entries: make([]TranscriptEntry, 0),
	}
}

func (t *Transcript) Add(e TranscriptEntry) {
	t.Entries = append(t.Entries, e)
}

func (t *Transcript) GetRMSTranscript() rms.VObjectTranscriptReport_Transcript {
	objectTranscript := rms.VObjectTranscriptReport_Transcript{}
	for _, entry := range t.Entries {
		rmsEntry := new(rms.Any)

		switch typedEntry := interface{}(entry).(type) {
		case TranscriptEntryIncomingRequest:
			rmsEntry.Set(
				&rms.VObjectTranscriptReport_TranscriptEntryIncomingRequest{
					ObjectMemory: rms.NewReference(typedEntry.ObjectMemory),
					Incoming:     rms.NewReference(typedEntry.Incoming),
					Request:      []byte{}, // fixme: waiting for messages to move to RMS package
				})
		case TranscriptEntryIncomingResult:
			rmsEntry.Set(&rms.VObjectTranscriptReport_TranscriptEntryIncomingResult{
				IncomingResult: rms.NewReference(typedEntry.IncomingResult),
				ObjectState: rms.NewReference(typedEntry.ObjectMemory),
			})

			// TODO add this later
			// case TranscriptEntryOutgoingRequest:
			// 	rmsEntry.Set(&rms.VObjectTranscriptReport_TranscriptEntryOutgoingRequest{})
			// case TranscriptEntryOutgoingResult:
			// 	rmsEntry.Set(&rms.VObjectTranscriptReport_TranscriptEntryOutgoingResult{})
		}

		objectTranscript.Entries = append(objectTranscript.Entries, *rmsEntry)
	}

	return objectTranscript
}
