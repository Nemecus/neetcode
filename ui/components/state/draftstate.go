package state

type draftState struct {
	document *document
}

func (ds *draftState) publish() {
	ds.document.setState(&reviewState{ds.document})
}
