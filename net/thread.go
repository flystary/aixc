package net

import (
	"context"
	"sync"

	"aixc/model"
)

type result struct {
	mode string
	ok   bool
}

func ThreadQueryMode(sn string) string {
	const worker = 4

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan result, worker)

	var wg sync.WaitGroup
	wg.Add(worker)

	for i := 1; i <= worker; i++ {
		mode := model.M(i).Enum()

		go func(m string) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			ok := SyncDataMemorybySnMode(sn, m)
			select {
			case <-ctx.Done():
				return
			case ch <- result{
				mode: m, ok: ok,
			}:
			}
		}(mode)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var ms []string

	for r := range ch {
		if !r.ok {
			continue
		}
		if r.mode == "valor" {
			cancel()
			return "valor"
		}
		ms = append(ms, r.mode)
	}

	return resolve(ms)
}

func resolve(ms []string) string {
	if len(ms) >= 2 {
		for _, m := range ms {
			if m == "yifeng" {
				return m
			}
		}
		for _, m := range ms {
			if m == "valor" {
				return m
			}
		}
		return "unknown"
	}

	if len(ms) == 1 {
		return ms[0]
	}

	return "unknown"
}
