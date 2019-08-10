package M2N

import "TuriteaWebResources/asynchronousIO"

func NewM2NMachine(dataSources []asynchronousIO.DataSource, beanTypeNumbers, exceptBufferNumber int64, reIn bool) asynchronousIO.AsynchronousIOMachine {
	if beanTypeNumbers == 1 {
		return newForOneDataSource(dataSources[0], beanTypeNumbers, exceptBufferNumber)
	}
	if reIn {
		return newReIn(dataSources)
	}
	return newNormal(dataSources, beanTypeNumbers, exceptBufferNumber)
}
