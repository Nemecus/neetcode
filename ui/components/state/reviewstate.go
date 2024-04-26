package state

type reviewState struct {
	document *document
}

func (rs *reviewState) publish() {
	if rs.document.isApproved() {
		rs.document.setState(&publishedState{document: rs.document})
	} else {
		rs.document.setState(&draftState{document: rs.document})
	}
}
