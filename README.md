# prefab
A tool to get prefabricated production ready code as a starter for your next adventure.

## Get started

```
> git clone https://github.com/yantrashala/prefab.git
...
> cd prefab
> docker build -t ps/fab .
...
> docker run --rm -it ps/fab ./fab
```

or if you have go installed locally try
```
> git clone https://github.com/yantrashala/prefab.git
...
> cd prefab
> go get ./
> go run main.go 
```


### Prerequisites
* [git client](https://git-scm.com/)
* [Docker 17.05 or later](https://www.docker.com/)
* [go 1.11 or later](https://golang.org/)

## Contributing
1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/prefab && cd prefab)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request