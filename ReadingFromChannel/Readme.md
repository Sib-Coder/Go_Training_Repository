# Ассинхронное чтение данных из каналов
```
	select {
		case msg := <-massage1:
			fmt.Println(msg)
		case msg := <-massage2:
			fmt.Println(msg)
		}
```