#### Generate the input file

Run `make generate`

#### Get the topN numbers

Run `go run main.go --numbers <the top n numbers to print> --file <input file>`

#### Code structure

```
├── algorithm
│   ├── heap.go // contains the min heap implementation
│   └── queue.go // contains the backing queue implementation for the heap
├── hack
│   └── generate.go // generates an input file with random integers on every rile
├── main.go // reads the file, and calls performs operations on the heap for every file
├── Makefile // utils to run the code
```
