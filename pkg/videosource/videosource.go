package videosource

type VS struct {
	id   string
	done chan bool
}

func (vs *VS) Process(sourceURL string) {

	for {
		select {
		case <-vs.done:
			{
				break
			}
		default:
			{
				// read image
				// process image
				// send results
			}
		}
	}

}
