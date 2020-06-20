module producer

go 1.13

replace github.com/rickihastings/go-redis-streams => ../../

require (
	github.com/go-redis/redis/v8 v8.0.0-beta.5
	github.com/rickihastings/go-redis-streams v0.0.0-20200620100836-5ad49841a597
	github.com/vmihailenco/msgpack v4.0.4+incompatible
)
