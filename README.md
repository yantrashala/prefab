# prefab
A tool to get prefabricated production ready code as a starter for your next adventure.

## Developers getting started

```
> git clone https://github.com/yantrashala/prefab.git
```

Build and run in docker (suggested)
```
> cd prefab
> docker build -t ps/fab .
...
> docker run --rm -it -p9876:9876 ps/fab ./fab server
```

or if you have prereqs installed locally try
```
> git clone https://github.com/yantrashala/prefab.git
...
> cd prefab
> go get ./
> cd ui
> npm install
> npm run build
> cd ..
> go run main.go 
```

or use the make file locally
```
> git clone https://github.com/yantrashala/prefab.git
...
> cd prefab
> make install
> make compile
> make start-server
...
...
> make stop-server
```


### Prerequisites
* [git client](https://git-scm.com/)
* [Docker 17.05 or later](https://www.docker.com/)
* [go 1.11 or later](https://golang.org/)
* [node v11 or later](https://nodejs.org)

## Contributing
1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/prefab && cd prefab)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request