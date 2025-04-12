package fan_in

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

// genChannel создает канал, в который последовательно отправляются элементы из среза values с указанной задержкой, затем канал закрывается.
func genChannel(values []int, delay time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			time.Sleep(delay)
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// TestMergeChannels_MultipleChannels проверяет, что функция корректно объединяет значения из нескольких входных каналов.
func TestMergeChannels_MultipleChannels(t *testing.T) {
	// Создаем два канала с разными значениями и задержками.
	ch1 := genChannel([]int{1, 3, 5}, 10*time.Millisecond)
	ch2 := genChannel([]int{2, 4, 6}, 15*time.Millisecond)

	merged := MergeChannels(ch1, ch2)
	if merged == nil {
		t.Errorf("Ожидалось, что выходной канал не будет nil")
		return
	}

	var results []int
	for val := range merged {
		results = append(results, val)
	}

	// Так как порядок может быть неопределенным, сортируем полученные значения для сравнения.
	expected := []int{1, 2, 3, 4, 5, 6}
	sort.Ints(results)
	sort.Ints(expected)
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, results)
	}
}

// TestMergeChannels_NoChannels проверяет, что функция корректно работает, когда входных каналов нет.
func TestMergeChannels_NoChannels(t *testing.T) {
	merged := MergeChannels()
	if merged == nil {
		t.Errorf("Ожидалось, что выходной канал не будет nil")
		return
	}

	select {
	case _, ok := <-merged:
		if ok {
			t.Error("Ожидалось, что канал сразу будет закрыт")
		}
	case <-time.After(50 * time.Millisecond):
		t.Error("Ожидалось, что канал сразу закроется")
	}
}

// TestMergeChannels_ChannelClose проверяет, что выходной канал закрывается после обработки всех входных каналов.
func TestMergeChannels_ChannelClose(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
		close(ch1)
	}()
	go func() {
		ch2 <- 20
		close(ch2)
	}()

	merged := MergeChannels(ch1, ch2)
	if merged == nil {
		t.Errorf("Ожидалось, что выходной канал не будет nil")
		return
	}

	var results []int
	for val := range merged {
		results = append(results, val)
	}
	expected := []int{10, 20}
	sort.Ints(results)
	sort.Ints(expected)
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, results)
	}
}
