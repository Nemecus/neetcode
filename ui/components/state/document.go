package state

type document struct {
	draft     state
	review    state
	publushed state

	currentState state
	approved     bool
}

func newDocument() *document {
	d := &document{
		approved: false,
	}

	draftState := &draftState{
		document: d,
	}

	reviewState := &reviewState{
		document: d,
	}

	publishedState := &publishedState{
		document: d,
	}

	d.setState(draftState)
	d.draft = draftState
	d.review = reviewState
	d.publushed = publishedState

	return d
}

func (d *document) publish() {
	d.currentState.publish()
}

func (d *document) isApproved() bool {
	return d.approved
}

func (d *document) setApproval(status bool) {
	d.approved = status
}

func (d *document) getState() state {
	return d.currentState
}

func (d *document) setState(s state) {
	d.currentState = s
}
