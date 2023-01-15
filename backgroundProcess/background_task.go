package backgroundProcess

func RunBackgroundTask() {
	go fetchingLatestPrices()
}
