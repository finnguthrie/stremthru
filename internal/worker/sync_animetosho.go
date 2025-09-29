package worker

import (
	"github.com/MunifTanjim/stremthru/internal/animetosho"
)

func InitSyncAnimeToshoWorker(conf *WorkerConfig) *Worker {
	conf.Executor = func(w *Worker) error {
		err := animetosho.SyncDataset()
		if err != nil {
			return err
		}
		return nil
	}

	worker := NewWorker(conf)
	return worker
}
