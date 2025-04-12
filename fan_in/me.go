package fan_in

// func MergeChannels(channels ...<-chan int) <-chan int {
// 	// Создаем выходной канал для передачи значений
// 	out := make(chan int)
//
// 	var wg sync.WaitGroup
//
// 	// Функция для обработки каждого входного канала
// 	// Читает все значения из входного канала и отправляет их в канал out.
// 	output := func(ch <-chan int, workerID int) {
// 		defer wg.Done()
// 		for val := range ch {
// 			// Можно добавить логирование, если необходимо: какой worker (workerID) обрабатывает значение
// 			out <- val
// 		}
// 	}
//
// 	// Устанавливаем количество горутин, равное числу входных каналов
// 	wg.Add(len(channels))
// 	for id, ch := range channels {
// 		go output(ch, id)
// 	}
//
// 	// Закрываем выходной канал, когда все входные каналы обработаны
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
//
// 	return out
// }
