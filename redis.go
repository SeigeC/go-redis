package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Config for redis config
type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

type Client struct {
	*redis.Client
	ctx    context.Context
	useCtx bool
}

func (c Client) Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if c.useCtx {
		return c.Client.Pipelined(c.ctx, fn)
	}
	ctx:= context.Background()
	return c.Client.Pipelined(ctx,fn)
}

func (c Client) TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if c.useCtx {
		return c.Client.TxPipelined(c.ctx, fn)
	}
	ctx := context.Background()
	return c.Client.TxPipelined(ctx, fn)
}

func (c Client) Command() *redis.CommandsInfoCmd {
	if c.useCtx {
		return c.Client.Command(c.ctx)
	}
	ctx := context.Background()
	return c.Client.Command(ctx)
}

func (c Client) ClientGetName() *redis.StringCmd {
	if c.useCtx {
		return c.Client.ClientGetName(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClientGetName(ctx)
}

func (c Client) Echo(message interface{}) *redis.StringCmd {
	if c.useCtx {
		return c.Client.Echo(c.ctx, message)
	}
	ctx := context.Background()
	return c.Client.Echo(ctx, message)
}

func (c Client) Ping() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Ping(c.ctx)
	}
	ctx := context.Background()
	return c.Client.Ping(ctx)
}

func (c Client) Quit() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Quit(c.ctx)
	}
	ctx := context.Background()
	return c.Client.Quit(ctx)
}

func (c Client) Del(keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Del(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.Del(ctx, keys...)
}

func (c Client) Unlink(keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Unlink(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.Unlink(ctx, keys...)
}

func (c Client) Dump(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.Dump(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.Dump(ctx, key)
}

func (c Client) Exists(keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Exists(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.Exists(ctx, keys...)
}

func (c Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.Expire(c.ctx, key, expiration)
	}
	ctx := context.Background()
	return c.Client.Expire(ctx, key, expiration)
}

func (c Client) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.ExpireAt(c.ctx, key, tm)
	}
	ctx := context.Background()
	return c.Client.ExpireAt(ctx, key, tm)
}

func (c Client) Keys(pattern string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.Keys(c.ctx, pattern)
	}
	ctx := context.Background()
	return c.Client.Keys(ctx, pattern)
}

func (c Client) Migrate(host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Migrate(c.ctx, host, port, key, db, timeout)
	}
	ctx := context.Background()
	return c.Client.Migrate(ctx, host, port, key, db, timeout)
}

func (c Client) Move(key string, db int) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.Move(c.ctx, key, db)
	}
	ctx := context.Background()
	return c.Client.Move(ctx, key, db)
}

func (c Client) ObjectRefCount(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ObjectRefCount(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.ObjectRefCount(ctx, key)
}

func (c Client) ObjectEncoding(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.ObjectEncoding(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.ObjectEncoding(ctx, key)
}

func (c Client) ObjectIdleTime(key string) *redis.DurationCmd {
	if c.useCtx {
		return c.Client.ObjectIdleTime(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.ObjectIdleTime(ctx, key)
}

func (c Client) Persist(key string) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.Persist(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.Persist(ctx, key)
}

func (c Client) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.PExpire(c.ctx, key, expiration)
	}
	ctx := context.Background()
	return c.Client.PExpire(ctx, key, expiration)
}

func (c Client) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.PExpireAt(c.ctx, key, tm)
	}
	ctx := context.Background()
	return c.Client.PExpireAt(ctx, key, tm)
}

func (c Client) PTTL(key string) *redis.DurationCmd {
	if c.useCtx {
		return c.Client.PTTL(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.PTTL(ctx, key)
}

func (c Client) RandomKey() *redis.StringCmd {
	if c.useCtx {
		return c.Client.RandomKey(c.ctx)
	}
	ctx := context.Background()
	return c.Client.RandomKey(ctx)
}

func (c Client) Rename(key, newkey string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Rename(c.ctx, key, newkey)
	}
	ctx := context.Background()
	return c.Client.Rename(ctx, key, newkey)
}

func (c Client) RenameNX(key, newkey string) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.RenameNX(c.ctx, key, newkey)
	}
	ctx := context.Background()
	return c.Client.RenameNX(ctx, key, newkey)
}

func (c Client) Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Restore(c.ctx, key, ttl, value)
	}
	ctx := context.Background()
	return c.Client.Restore(ctx, key, ttl, value)
}

func (c Client) RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.RestoreReplace(c.ctx, key, ttl, value)
	}
	ctx := context.Background()
	return c.Client.RestoreReplace(ctx, key, ttl, value)
}

func (c Client) Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.Sort(c.ctx, key, sort)
	}
	ctx := context.Background()
	return c.Client.Sort(ctx, key, sort)
}

func (c Client) SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SortStore(c.ctx, key, store, sort)
	}
	ctx := context.Background()
	return c.Client.SortStore(ctx, key, store, sort)
}

func (c Client) SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	if c.useCtx {
		return c.Client.SortInterfaces(c.ctx, key, sort)
	}
	ctx := context.Background()
	return c.Client.SortInterfaces(ctx, key, sort)
}

func (c Client) Touch(keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Touch(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.Touch(ctx, keys...)
}

func (c Client) TTL(key string) *redis.DurationCmd {
	if c.useCtx {
		return c.Client.TTL(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.TTL(ctx, key)
}

func (c Client) Type(key string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Type(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.Type(ctx, key)
}

func (c Client) Append(key, value string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Append(c.ctx, key, value)
	}
	ctx := context.Background()
	return c.Client.Append(ctx, key, value)
}

func (c Client) Decr(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Decr(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.Decr(ctx, key)
}

func (c Client) DecrBy(key string, decrement int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.DecrBy(c.ctx, key, decrement)
	}
	ctx := context.Background()
	return c.Client.DecrBy(ctx, key, decrement)
}

func (c Client) Get(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.Get(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.Get(ctx, key)
}

func (c Client) GetRange(key string, start, end int64) *redis.StringCmd {
	if c.useCtx {
		return c.Client.GetRange(c.ctx, key, start, end)
	}
	ctx := context.Background()
	return c.Client.GetRange(ctx, key, start, end)
}

func (c Client) GetSet(key string, value interface{}) *redis.StringCmd {
	if c.useCtx {
		return c.Client.GetSet(c.ctx, key, value)
	}
	ctx := context.Background()
	return c.Client.GetSet(ctx, key, value)
}

func (c Client) GetEx(key string, expiration time.Duration) *redis.StringCmd {
	if c.useCtx {
		return c.Client.GetEx(c.ctx, key, expiration)
	}
	ctx := context.Background()
	return c.Client.GetEx(ctx, key, expiration)
}

func (c Client) GetDel(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.GetDel(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.GetDel(ctx, key)
}

func (c Client) Incr(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Incr(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.Incr(ctx, key)
}

func (c Client) IncrBy(key string, value int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.IncrBy(c.ctx, key, value)
	}
	ctx := context.Background()
	return c.Client.IncrBy(ctx, key, value)
}

func (c Client) IncrByFloat(key string, value float64) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.IncrByFloat(c.ctx, key, value)
	}
	ctx := context.Background()
	return c.Client.IncrByFloat(ctx, key, value)
}

func (c Client) MGet(keys ...string) *redis.SliceCmd {
	if c.useCtx {
		return c.Client.MGet(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.MGet(ctx, keys...)
}

func (c Client) MSet(values ...interface{}) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.MSet(c.ctx, values...)
	}
	ctx := context.Background()
	return c.Client.MSet(ctx, values...)
}

func (c Client) MSetNX(values ...interface{}) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.MSetNX(c.ctx, values...)
	}
	ctx := context.Background()
	return c.Client.MSetNX(ctx, values...)
}

func (c Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Set(c.ctx, key, value, expiration)
	}
	ctx := context.Background()
	return c.Client.Set(ctx, key, value, expiration)
}

func (c Client) SetArgs(key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.SetArgs(c.ctx, key, value, a)
	}
	ctx := context.Background()
	return c.Client.SetArgs(ctx, key, value, a)
}

func (c Client) SetEX(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.SetEX(c.ctx, key, value, expiration)
	}
	ctx := context.Background()
	return c.Client.SetEX(ctx, key, value, expiration)
}

func (c Client) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.SetNX(c.ctx, key, value, expiration)
	}
	ctx := context.Background()
	return c.Client.SetNX(ctx, key, value, expiration)
}

func (c Client) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.SetXX(c.ctx, key, value, expiration)
	}
	ctx := context.Background()
	return c.Client.SetXX(ctx, key, value, expiration)
}

func (c Client) SetRange(key string, offset int64, value string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SetRange(c.ctx, key, offset, value)
	}
	ctx := context.Background()
	return c.Client.SetRange(ctx, key, offset, value)
}

func (c Client) StrLen(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.StrLen(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.StrLen(ctx, key)
}

func (c Client) GetBit(key string, offset int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.GetBit(c.ctx, key, offset)
	}
	ctx := context.Background()
	return c.Client.GetBit(ctx, key, offset)
}

func (c Client) SetBit(key string, offset int64, value int) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SetBit(c.ctx, key, offset, value)
	}
	ctx := context.Background()
	return c.Client.SetBit(ctx, key, offset, value)
}

func (c Client) BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	if c.useCtx {
		return c.Client.BitCount(c.ctx, key, bitCount)
	}
	ctx := context.Background()
	return c.Client.BitCount(ctx, key, bitCount)
}

func (c Client) BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.BitOpAnd(c.ctx, destKey, keys...)
	}
	ctx := context.Background()
	return c.Client.BitOpAnd(ctx, destKey, keys...)
}

func (c Client) BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.BitOpOr(c.ctx, destKey, keys...)
	}
	ctx := context.Background()
	return c.Client.BitOpOr(ctx, destKey, keys...)
}

func (c Client) BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.BitOpXor(c.ctx, destKey, keys...)
	}
	ctx := context.Background()
	return c.Client.BitOpXor(ctx, destKey, keys...)
}

func (c Client) BitOpNot(destKey string, key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.BitOpNot(c.ctx, destKey, key)
	}
	ctx := context.Background()
	return c.Client.BitOpNot(ctx, destKey, key)
}

func (c Client) BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.BitPos(c.ctx, key, bit, pos...)
	}
	ctx := context.Background()
	return c.Client.BitPos(ctx, key, bit, pos...)
}

func (c Client) BitField(key string, args ...interface{}) *redis.IntSliceCmd {
	if c.useCtx {
		return c.Client.BitField(c.ctx, key, args...)
	}
	ctx := context.Background()
	return c.Client.BitField(ctx, key, args...)
}

func (c Client) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	if c.useCtx {
		return c.Client.Scan(c.ctx, cursor, match, count)
	}
	ctx := context.Background()
	return c.Client.Scan(ctx, cursor, match, count)
}

func (c Client) ScanType(cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	if c.useCtx {
		return c.Client.ScanType(c.ctx, cursor, match, count, keyType)
	}
	ctx := context.Background()
	return c.Client.ScanType(ctx, cursor, match, count, keyType)
}

func (c Client) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if c.useCtx {
		return c.Client.SScan(c.ctx, key, cursor, match, count)
	}
	ctx := context.Background()
	return c.Client.SScan(ctx, key, cursor, match, count)
}

func (c Client) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if c.useCtx {
		return c.Client.HScan(c.ctx, key, cursor, match, count)
	}
	ctx := context.Background()
	return c.Client.HScan(ctx, key, cursor, match, count)
}

func (c Client) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if c.useCtx {
		return c.Client.ZScan(c.ctx, key, cursor, match, count)
	}
	ctx := context.Background()
	return c.Client.ZScan(ctx, key, cursor, match, count)
}

func (c Client) HDel(key string, fields ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.HDel(c.ctx, key, fields...)
	}
	ctx := context.Background()
	return c.Client.HDel(ctx, key, fields...)
}

func (c Client) HExists(key, field string) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.HExists(c.ctx, key, field)
	}
	ctx := context.Background()
	return c.Client.HExists(ctx, key, field)
}

func (c Client) HGet(key, field string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.HGet(c.ctx, key, field)
	}
	ctx := context.Background()
	return c.Client.HGet(ctx, key, field)
}

func (c Client) HGetAll(key string) *redis.StringStringMapCmd {
	if c.useCtx {
		return c.Client.HGetAll(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.HGetAll(ctx, key)
}

func (c Client) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.HIncrBy(c.ctx, key, field, incr)
	}
	ctx := context.Background()
	return c.Client.HIncrBy(ctx, key, field, incr)
}

func (c Client) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.HIncrByFloat(c.ctx, key, field, incr)
	}
	ctx := context.Background()
	return c.Client.HIncrByFloat(ctx, key, field, incr)
}

func (c Client) HKeys(key string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.HKeys(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.HKeys(ctx, key)
}

func (c Client) HLen(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.HLen(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.HLen(ctx, key)
}

func (c Client) HMGet(key string, fields ...string) *redis.SliceCmd {
	if c.useCtx {
		return c.Client.HMGet(c.ctx, key, fields...)
	}
	ctx := context.Background()
	return c.Client.HMGet(ctx, key, fields...)
}

func (c Client) HSet(key string, values ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.HSet(c.ctx, key, values...)
	}
	ctx := context.Background()
	return c.Client.HSet(ctx, key, values...)
}

func (c Client) HMSet(key string, values ...interface{}) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.HMSet(c.ctx, key, values...)
	}
	ctx := context.Background()
	return c.Client.HMSet(ctx, key, values...)
}

func (c Client) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.HSetNX(c.ctx, key, field, value)
	}
	ctx := context.Background()
	return c.Client.HSetNX(ctx, key, field, value)
}

func (c Client) HVals(key string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.HVals(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.HVals(ctx, key)
}

func (c Client) HRandField(key string, count int, withValues bool) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.HRandField(c.ctx, key, count, withValues)
	}
	ctx := context.Background()
	return c.Client.HRandField(ctx, key, count, withValues)
}

func (c Client) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.BLPop(c.ctx, timeout, keys...)
	}
	ctx := context.Background()
	return c.Client.BLPop(ctx, timeout, keys...)
}

func (c Client) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.BRPop(c.ctx, timeout, keys...)
	}
	ctx := context.Background()
	return c.Client.BRPop(ctx, timeout, keys...)
}

func (c Client) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	if c.useCtx {
		return c.Client.BRPopLPush(c.ctx, source, destination, timeout)
	}
	ctx := context.Background()
	return c.Client.BRPopLPush(ctx, source, destination, timeout)
}

func (c Client) LIndex(key string, index int64) *redis.StringCmd {
	if c.useCtx {
		return c.Client.LIndex(c.ctx, key, index)
	}
	ctx := context.Background()
	return c.Client.LIndex(ctx, key, index)
}

func (c Client) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LInsert(c.ctx, key, op, pivot, value)
	}
	ctx := context.Background()
	return c.Client.LInsert(ctx, key, op, pivot, value)
}

func (c Client) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LInsertBefore(c.ctx, key, pivot, value)
	}
	ctx := context.Background()
	return c.Client.LInsertBefore(ctx, key, pivot, value)
}

func (c Client) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LInsertAfter(c.ctx, key, pivot, value)
	}
	ctx := context.Background()
	return c.Client.LInsertAfter(ctx, key, pivot, value)
}

func (c Client) LLen(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LLen(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.LLen(ctx, key)
}

func (c Client) LPop(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.LPop(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.LPop(ctx, key)
}

func (c Client) LPopCount(key string, count int) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.LPopCount(c.ctx, key, count)
	}
	ctx := context.Background()
	return c.Client.LPopCount(ctx, key, count)
}

func (c Client) LPos(key string, value string, args redis.LPosArgs) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LPos(c.ctx, key, value, args)
	}
	ctx := context.Background()
	return c.Client.LPos(ctx, key, value, args)
}

func (c Client) LPosCount(key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	if c.useCtx {
		return c.Client.LPosCount(c.ctx, key, value, count, args)
	}
	ctx := context.Background()
	return c.Client.LPosCount(ctx, key, value, count, args)
}

func (c Client) LPush(key string, values ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LPush(c.ctx, key, values...)
	}
	ctx := context.Background()
	return c.Client.LPush(ctx, key, values...)
}

func (c Client) LPushX(key string, values ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LPushX(c.ctx, key, values...)
	}
	ctx := context.Background()
	return c.Client.LPushX(ctx, key, values...)
}

func (c Client) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.LRange(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.LRange(ctx, key, start, stop)
}

func (c Client) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.LRem(c.ctx, key, count, value)
	}
	ctx := context.Background()
	return c.Client.LRem(ctx, key, count, value)
}

func (c Client) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.LSet(c.ctx, key, index, value)
	}
	ctx := context.Background()
	return c.Client.LSet(ctx, key, index, value)
}

func (c Client) LTrim(key string, start, stop int64) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.LTrim(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.LTrim(ctx, key, start, stop)
}

func (c Client) RPop(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.RPop(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.RPop(ctx, key)
}

func (c Client) RPopCount(key string, count int) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.RPopCount(c.ctx, key, count)
	}
	ctx := context.Background()
	return c.Client.RPopCount(ctx, key, count)
}

func (c Client) RPopLPush(source, destination string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.RPopLPush(c.ctx, source, destination)
	}
	ctx := context.Background()
	return c.Client.RPopLPush(ctx, source, destination)
}

func (c Client) RPush(key string, values ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.RPush(c.ctx, key, values...)
	}
	ctx := context.Background()
	return c.Client.RPush(ctx, key, values...)
}

func (c Client) RPushX(key string, values ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.RPushX(c.ctx, key, values...)
	}
	ctx := context.Background()
	return c.Client.RPushX(ctx, key, values...)
}

func (c Client) LMove(source, destination, srcpos, destpos string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.LMove(c.ctx, source, destination, srcpos, destpos)
	}
	ctx := context.Background()
	return c.Client.LMove(ctx, source, destination, srcpos, destpos)
}

func (c Client) BLMove(source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	if c.useCtx {
		return c.Client.BLMove(c.ctx, source, destination, srcpos, destpos, timeout)
	}
	ctx := context.Background()
	return c.Client.BLMove(ctx, source, destination, srcpos, destpos, timeout)
}

func (c Client) SAdd(key string, members ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SAdd(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.SAdd(ctx, key, members...)
}

func (c Client) SCard(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SCard(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.SCard(ctx, key)
}

func (c Client) SDiff(keys ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.SDiff(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.SDiff(ctx, keys...)
}

func (c Client) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SDiffStore(c.ctx, destination, keys...)
	}
	ctx := context.Background()
	return c.Client.SDiffStore(ctx, destination, keys...)
}

func (c Client) SInter(keys ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.SInter(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.SInter(ctx, keys...)
}

func (c Client) SInterStore(destination string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SInterStore(c.ctx, destination, keys...)
	}
	ctx := context.Background()
	return c.Client.SInterStore(ctx, destination, keys...)
}

func (c Client) SIsMember(key string, member interface{}) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.SIsMember(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.SIsMember(ctx, key, member)
}

func (c Client) SMIsMember(key string, members ...interface{}) *redis.BoolSliceCmd {
	if c.useCtx {
		return c.Client.SMIsMember(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.SMIsMember(ctx, key, members...)
}

func (c Client) SMembers(key string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.SMembers(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.SMembers(ctx, key)
}

func (c Client) SMembersMap(key string) *redis.StringStructMapCmd {
	if c.useCtx {
		return c.Client.SMembersMap(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.SMembersMap(ctx, key)
}

func (c Client) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.SMove(c.ctx, source, destination, member)
	}
	ctx := context.Background()
	return c.Client.SMove(ctx, source, destination, member)
}

func (c Client) SPop(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.SPop(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.SPop(ctx, key)
}

func (c Client) SPopN(key string, count int64) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.SPopN(c.ctx, key, count)
	}
	ctx := context.Background()
	return c.Client.SPopN(ctx, key, count)
}

func (c Client) SRandMember(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.SRandMember(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.SRandMember(ctx, key)
}

func (c Client) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.SRandMemberN(c.ctx, key, count)
	}
	ctx := context.Background()
	return c.Client.SRandMemberN(ctx, key, count)
}

func (c Client) SRem(key string, members ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SRem(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.SRem(ctx, key, members...)
}

func (c Client) SUnion(keys ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.SUnion(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.SUnion(ctx, keys...)
}

func (c Client) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.SUnionStore(c.ctx, destination, keys...)
	}
	ctx := context.Background()
	return c.Client.SUnionStore(ctx, destination, keys...)
}

func (c Client) XAdd(a *redis.XAddArgs) *redis.StringCmd {
	if c.useCtx {
		return c.Client.XAdd(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XAdd(ctx, a)
}

func (c Client) XDel(stream string, ids ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XDel(c.ctx, stream, ids...)
	}
	ctx := context.Background()
	return c.Client.XDel(ctx, stream, ids...)
}

func (c Client) XLen(stream string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XLen(c.ctx, stream)
	}
	ctx := context.Background()
	return c.Client.XLen(ctx, stream)
}

func (c Client) XRange(stream, start, stop string) *redis.XMessageSliceCmd {
	if c.useCtx {
		return c.Client.XRange(c.ctx, stream, start, stop)
	}
	ctx := context.Background()
	return c.Client.XRange(ctx, stream, start, stop)
}

func (c Client) XRangeN(stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	if c.useCtx {
		return c.Client.XRangeN(c.ctx, stream, start, stop, count)
	}
	ctx := context.Background()
	return c.Client.XRangeN(ctx, stream, start, stop, count)
}

func (c Client) XRevRange(stream string, start, stop string) *redis.XMessageSliceCmd {
	if c.useCtx {
		return c.Client.XRevRange(c.ctx, stream, start, stop)
	}
	ctx := context.Background()
	return c.Client.XRevRange(ctx, stream, start, stop)
}

func (c Client) XRevRangeN(stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	if c.useCtx {
		return c.Client.XRevRangeN(c.ctx, stream, start, stop, count)
	}
	ctx := context.Background()
	return c.Client.XRevRangeN(ctx, stream, start, stop, count)
}

func (c Client) XRead(a *redis.XReadArgs) *redis.XStreamSliceCmd {
	if c.useCtx {
		return c.Client.XRead(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XRead(ctx, a)
}

func (c Client) XReadStreams(streams ...string) *redis.XStreamSliceCmd {
	if c.useCtx {
		return c.Client.XReadStreams(c.ctx, streams...)
	}
	ctx := context.Background()
	return c.Client.XReadStreams(ctx, streams...)
}

func (c Client) XGroupCreate(stream, group, start string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.XGroupCreate(c.ctx, stream, group, start)
	}
	ctx := context.Background()
	return c.Client.XGroupCreate(ctx, stream, group, start)
}

func (c Client) XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.XGroupCreateMkStream(c.ctx, stream, group, start)
	}
	ctx := context.Background()
	return c.Client.XGroupCreateMkStream(ctx, stream, group, start)
}

func (c Client) XGroupSetID(stream, group, start string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.XGroupSetID(c.ctx, stream, group, start)
	}
	ctx := context.Background()
	return c.Client.XGroupSetID(ctx, stream, group, start)
}

func (c Client) XGroupDestroy(stream, group string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XGroupDestroy(c.ctx, stream, group)
	}
	ctx := context.Background()
	return c.Client.XGroupDestroy(ctx, stream, group)
}

func (c Client) XGroupCreateConsumer(stream, group, consumer string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XGroupCreateConsumer(c.ctx, stream, group, consumer)
	}
	ctx := context.Background()
	return c.Client.XGroupCreateConsumer(ctx, stream, group, consumer)
}

func (c Client) XGroupDelConsumer(stream, group, consumer string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XGroupDelConsumer(c.ctx, stream, group, consumer)
	}
	ctx := context.Background()
	return c.Client.XGroupDelConsumer(ctx, stream, group, consumer)
}

func (c Client) XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	if c.useCtx {
		return c.Client.XReadGroup(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XReadGroup(ctx, a)
}

func (c Client) XAck(stream, group string, ids ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XAck(c.ctx, stream, group, ids...)
	}
	ctx := context.Background()
	return c.Client.XAck(ctx, stream, group, ids...)
}

func (c Client) XPending(stream, group string) *redis.XPendingCmd {
	if c.useCtx {
		return c.Client.XPending(c.ctx, stream, group)
	}
	ctx := context.Background()
	return c.Client.XPending(ctx, stream, group)
}

func (c Client) XPendingExt(a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	if c.useCtx {
		return c.Client.XPendingExt(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XPendingExt(ctx, a)
}

func (c Client) XClaim(a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	if c.useCtx {
		return c.Client.XClaim(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XClaim(ctx, a)
}

func (c Client) XClaimJustID(a *redis.XClaimArgs) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.XClaimJustID(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XClaimJustID(ctx, a)
}

func (c Client) XAutoClaim(a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	if c.useCtx {
		return c.Client.XAutoClaim(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XAutoClaim(ctx, a)
}

func (c Client) XAutoClaimJustID(a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	if c.useCtx {
		return c.Client.XAutoClaimJustID(c.ctx, a)
	}
	ctx := context.Background()
	return c.Client.XAutoClaimJustID(ctx, a)
}

func (c Client) XTrim(key string, maxLen int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XTrim(c.ctx, key, maxLen)
	}
	ctx := context.Background()
	return c.Client.XTrim(ctx, key, maxLen)
}

func (c Client) XTrimApprox(key string, maxLen int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XTrimApprox(c.ctx, key, maxLen)
	}
	ctx := context.Background()
	return c.Client.XTrimApprox(ctx, key, maxLen)
}

func (c Client) XTrimMaxLen(key string, maxLen int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XTrimMaxLen(c.ctx, key, maxLen)
	}
	ctx := context.Background()
	return c.Client.XTrimMaxLen(ctx, key, maxLen)
}

func (c Client) XTrimMaxLenApprox(key string, maxLen, limit int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XTrimMaxLenApprox(c.ctx, key, maxLen, limit)
	}
	ctx := context.Background()
	return c.Client.XTrimMaxLenApprox(ctx, key, maxLen, limit)
}

func (c Client) XTrimMinID(key string, minID string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XTrimMinID(c.ctx, key, minID)
	}
	ctx := context.Background()
	return c.Client.XTrimMinID(ctx, key, minID)
}

func (c Client) XTrimMinIDApprox(key string, minID string, limit int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.XTrimMinIDApprox(c.ctx, key, minID, limit)
	}
	ctx := context.Background()
	return c.Client.XTrimMinIDApprox(ctx, key, minID, limit)
}

func (c Client) XInfoGroups(key string) *redis.XInfoGroupsCmd {
	if c.useCtx {
		return c.Client.XInfoGroups(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.XInfoGroups(ctx, key)
}

func (c Client) XInfoStream(key string) *redis.XInfoStreamCmd {
	if c.useCtx {
		return c.Client.XInfoStream(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.XInfoStream(ctx, key)
}

func (c Client) XInfoStreamFull(key string, count int) *redis.XInfoStreamFullCmd {
	if c.useCtx {
		return c.Client.XInfoStreamFull(c.ctx, key, count)
	}
	ctx := context.Background()
	return c.Client.XInfoStreamFull(ctx, key, count)
}

func (c Client) XInfoConsumers(key string, group string) *redis.XInfoConsumersCmd {
	if c.useCtx {
		return c.Client.XInfoConsumers(c.ctx, key, group)
	}
	ctx := context.Background()
	return c.Client.XInfoConsumers(ctx, key, group)
}

func (c Client) BZPopMax(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if c.useCtx {
		return c.Client.BZPopMax(c.ctx, timeout, keys...)
	}
	ctx := context.Background()
	return c.Client.BZPopMax(ctx, timeout, keys...)
}

func (c Client) BZPopMin(timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if c.useCtx {
		return c.Client.BZPopMin(c.ctx, timeout, keys...)
	}
	ctx := context.Background()
	return c.Client.BZPopMin(ctx, timeout, keys...)
}

func (c Client) ZAdd(key string, members ...*redis.Z) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAdd(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZAdd(ctx, key, members...)
}

func (c Client) ZAddNX(key string, members ...*redis.Z) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAddNX(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZAddNX(ctx, key, members...)
}

func (c Client) ZAddXX(key string, members ...*redis.Z) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAddXX(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZAddXX(ctx, key, members...)
}

func (c Client) ZAddCh(key string, members ...*redis.Z) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAddCh(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZAddCh(ctx, key, members...)
}

func (c Client) ZAddNXCh(key string, members ...*redis.Z) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAddNXCh(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZAddNXCh(ctx, key, members...)
}

func (c Client) ZAddXXCh(key string, members ...*redis.Z) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAddXXCh(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZAddXXCh(ctx, key, members...)
}

func (c Client) ZAddArgs(key string, args redis.ZAddArgs) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZAddArgs(c.ctx, key, args)
	}
	ctx := context.Background()
	return c.Client.ZAddArgs(ctx, key, args)
}

func (c Client) ZAddArgsIncr(key string, args redis.ZAddArgs) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.ZAddArgsIncr(c.ctx, key, args)
	}
	ctx := context.Background()
	return c.Client.ZAddArgsIncr(ctx, key, args)
}

func (c Client) ZIncr(key string, member *redis.Z) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.ZIncr(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.ZIncr(ctx, key, member)
}

func (c Client) ZIncrNX(key string, member *redis.Z) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.ZIncrNX(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.ZIncrNX(ctx, key, member)
}

func (c Client) ZIncrXX(key string, member *redis.Z) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.ZIncrXX(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.ZIncrXX(ctx, key, member)
}

func (c Client) ZCard(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZCard(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.ZCard(ctx, key)
}

func (c Client) ZCount(key, min, max string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZCount(c.ctx, key, min, max)
	}
	ctx := context.Background()
	return c.Client.ZCount(ctx, key, min, max)
}

func (c Client) ZLexCount(key, min, max string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZLexCount(c.ctx, key, min, max)
	}
	ctx := context.Background()
	return c.Client.ZLexCount(ctx, key, min, max)
}

func (c Client) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.ZIncrBy(c.ctx, key, increment, member)
	}
	ctx := context.Background()
	return c.Client.ZIncrBy(ctx, key, increment, member)
}

func (c Client) ZInter(store *redis.ZStore) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZInter(c.ctx, store)
	}
	ctx := context.Background()
	return c.Client.ZInter(ctx, store)
}

func (c Client) ZInterWithScores(store *redis.ZStore) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZInterWithScores(c.ctx, store)
	}
	ctx := context.Background()
	return c.Client.ZInterWithScores(ctx, store)
}

func (c Client) ZInterStore(destination string, store *redis.ZStore) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZInterStore(c.ctx, destination, store)
	}
	ctx := context.Background()
	return c.Client.ZInterStore(ctx, destination, store)
}

func (c Client) ZMScore(key string, members ...string) *redis.FloatSliceCmd {
	if c.useCtx {
		return c.Client.ZMScore(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZMScore(ctx, key, members...)
}

func (c Client) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZPopMax(c.ctx, key, count...)
	}
	ctx := context.Background()
	return c.Client.ZPopMax(ctx, key, count...)
}

func (c Client) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZPopMin(c.ctx, key, count...)
	}
	ctx := context.Background()
	return c.Client.ZPopMin(ctx, key, count...)
}

func (c Client) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRange(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.ZRange(ctx, key, start, stop)
}

func (c Client) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZRangeWithScores(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.ZRangeWithScores(ctx, key, start, stop)
}

func (c Client) ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRangeByScore(c.ctx, key, opt)
	}
	ctx := context.Background()
	return c.Client.ZRangeByScore(ctx, key, opt)
}

func (c Client) ZRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRangeByLex(c.ctx, key, opt)
	}
	ctx := context.Background()
	return c.Client.ZRangeByLex(ctx, key, opt)
}

func (c Client) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZRangeByScoreWithScores(c.ctx, key, opt)
	}
	ctx := context.Background()
	return c.Client.ZRangeByScoreWithScores(ctx, key, opt)
}

func (c Client) ZRangeArgs(z redis.ZRangeArgs) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRangeArgs(c.ctx, z)
	}
	ctx := context.Background()
	return c.Client.ZRangeArgs(ctx, z)
}

func (c Client) ZRangeArgsWithScores(z redis.ZRangeArgs) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZRangeArgsWithScores(c.ctx, z)
	}
	ctx := context.Background()
	return c.Client.ZRangeArgsWithScores(ctx, z)
}

func (c Client) ZRangeStore(dst string, z redis.ZRangeArgs) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRangeStore(c.ctx, dst, z)
	}
	ctx := context.Background()
	return c.Client.ZRangeStore(ctx, dst, z)
}

func (c Client) ZRank(key, member string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRank(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.ZRank(ctx, key, member)
}

func (c Client) ZRem(key string, members ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRem(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.ZRem(ctx, key, members...)
}

func (c Client) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRemRangeByRank(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.ZRemRangeByRank(ctx, key, start, stop)
}

func (c Client) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRemRangeByScore(c.ctx, key, min, max)
	}
	ctx := context.Background()
	return c.Client.ZRemRangeByScore(ctx, key, min, max)
}

func (c Client) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRemRangeByLex(c.ctx, key, min, max)
	}
	ctx := context.Background()
	return c.Client.ZRemRangeByLex(ctx, key, min, max)
}

func (c Client) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRevRange(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.ZRevRange(ctx, key, start, stop)
}

func (c Client) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZRevRangeWithScores(c.ctx, key, start, stop)
	}
	ctx := context.Background()
	return c.Client.ZRevRangeWithScores(ctx, key, start, stop)
}

func (c Client) ZRevRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRevRangeByScore(c.ctx, key, opt)
	}
	ctx := context.Background()
	return c.Client.ZRevRangeByScore(ctx, key, opt)
}

func (c Client) ZRevRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRevRangeByLex(c.ctx, key, opt)
	}
	ctx := context.Background()
	return c.Client.ZRevRangeByLex(ctx, key, opt)
}

func (c Client) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZRevRangeByScoreWithScores(c.ctx, key, opt)
	}
	ctx := context.Background()
	return c.Client.ZRevRangeByScoreWithScores(ctx, key, opt)
}

func (c Client) ZRevRank(key, member string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZRevRank(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.ZRevRank(ctx, key, member)
}

func (c Client) ZScore(key, member string) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.ZScore(c.ctx, key, member)
	}
	ctx := context.Background()
	return c.Client.ZScore(ctx, key, member)
}

func (c Client) ZUnionStore(dest string, store *redis.ZStore) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZUnionStore(c.ctx, dest, store)
	}
	ctx := context.Background()
	return c.Client.ZUnionStore(ctx, dest, store)
}

func (c Client) ZUnion(store redis.ZStore) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZUnion(c.ctx, store)
	}
	ctx := context.Background()
	return c.Client.ZUnion(ctx, store)
}

func (c Client) ZUnionWithScores(store redis.ZStore) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZUnionWithScores(c.ctx, store)
	}
	ctx := context.Background()
	return c.Client.ZUnionWithScores(ctx, store)
}

func (c Client) ZRandMember(key string, count int, withScores bool) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZRandMember(c.ctx, key, count, withScores)
	}
	ctx := context.Background()
	return c.Client.ZRandMember(ctx, key, count, withScores)
}

func (c Client) ZDiff(keys ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ZDiff(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.ZDiff(ctx, keys...)
}

func (c Client) ZDiffWithScores(keys ...string) *redis.ZSliceCmd {
	if c.useCtx {
		return c.Client.ZDiffWithScores(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.ZDiffWithScores(ctx, keys...)
}

func (c Client) ZDiffStore(destination string, keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ZDiffStore(c.ctx, destination, keys...)
	}
	ctx := context.Background()
	return c.Client.ZDiffStore(ctx, destination, keys...)
}

func (c Client) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.PFAdd(c.ctx, key, els...)
	}
	ctx := context.Background()
	return c.Client.PFAdd(ctx, key, els...)
}

func (c Client) PFCount(keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.PFCount(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.PFCount(ctx, keys...)
}

func (c Client) PFMerge(dest string, keys ...string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.PFMerge(c.ctx, dest, keys...)
	}
	ctx := context.Background()
	return c.Client.PFMerge(ctx, dest, keys...)
}

func (c Client) BgRewriteAOF() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.BgRewriteAOF(c.ctx)
	}
	ctx := context.Background()
	return c.Client.BgRewriteAOF(ctx)
}

func (c Client) BgSave() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.BgSave(c.ctx)
	}
	ctx := context.Background()
	return c.Client.BgSave(ctx)
}

func (c Client) ClientKill(ipPort string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClientKill(c.ctx, ipPort)
	}
	ctx := context.Background()
	return c.Client.ClientKill(ctx, ipPort)
}

func (c Client) ClientKillByFilter(keys ...string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ClientKillByFilter(c.ctx, keys...)
	}
	ctx := context.Background()
	return c.Client.ClientKillByFilter(ctx, keys...)
}

func (c Client) ClientList() *redis.StringCmd {
	if c.useCtx {
		return c.Client.ClientList(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClientList(ctx)
}

func (c Client) ClientPause(dur time.Duration) *redis.BoolCmd {
	if c.useCtx {
		return c.Client.ClientPause(c.ctx, dur)
	}
	ctx := context.Background()
	return c.Client.ClientPause(ctx, dur)
}

func (c Client) ClientID() *redis.IntCmd {
	if c.useCtx {
		return c.Client.ClientID(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClientID(ctx)
}

func (c Client) ConfigGet(parameter string) *redis.SliceCmd {
	if c.useCtx {
		return c.Client.ConfigGet(c.ctx, parameter)
	}
	ctx := context.Background()
	return c.Client.ConfigGet(ctx, parameter)
}

func (c Client) ConfigResetStat() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ConfigResetStat(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ConfigResetStat(ctx)
}

func (c Client) ConfigSet(parameter, value string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ConfigSet(c.ctx, parameter, value)
	}
	ctx := context.Background()
	return c.Client.ConfigSet(ctx, parameter, value)
}

func (c Client) ConfigRewrite() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ConfigRewrite(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ConfigRewrite(ctx)
}

func (c Client) DBSize() *redis.IntCmd {
	if c.useCtx {
		return c.Client.DBSize(c.ctx)
	}
	ctx := context.Background()
	return c.Client.DBSize(ctx)
}

func (c Client) FlushAll() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.FlushAll(c.ctx)
	}
	ctx := context.Background()
	return c.Client.FlushAll(ctx)
}

func (c Client) FlushAllAsync() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.FlushAllAsync(c.ctx)
	}
	ctx := context.Background()
	return c.Client.FlushAllAsync(ctx)
}

func (c Client) FlushDB() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.FlushDB(c.ctx)
	}
	ctx := context.Background()
	return c.Client.FlushDB(ctx)
}

func (c Client) FlushDBAsync() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.FlushDBAsync(c.ctx)
	}
	ctx := context.Background()
	return c.Client.FlushDBAsync(ctx)
}

func (c Client) Info(section ...string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.Info(c.ctx, section...)
	}
	ctx := context.Background()
	return c.Client.Info(ctx, section...)
}

func (c Client) LastSave() *redis.IntCmd {
	if c.useCtx {
		return c.Client.LastSave(c.ctx)
	}
	ctx := context.Background()
	return c.Client.LastSave(ctx)
}

func (c Client) Save() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Save(c.ctx)
	}
	ctx := context.Background()
	return c.Client.Save(ctx)
}

func (c Client) Shutdown() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.Shutdown(c.ctx)
	}
	ctx := context.Background()
	return c.Client.Shutdown(ctx)
}

func (c Client) ShutdownSave() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ShutdownSave(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ShutdownSave(ctx)
}

func (c Client) ShutdownNoSave() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ShutdownNoSave(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ShutdownNoSave(ctx)
}

func (c Client) SlaveOf(host, port string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.SlaveOf(c.ctx, host, port)
	}
	ctx := context.Background()
	return c.Client.SlaveOf(ctx, host, port)
}

func (c Client) Time() *redis.TimeCmd {
	if c.useCtx {
		return c.Client.Time(c.ctx)
	}
	ctx := context.Background()
	return c.Client.Time(ctx)
}

func (c Client) DebugObject(key string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.DebugObject(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.DebugObject(ctx, key)
}

func (c Client) ReadOnly() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ReadOnly(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ReadOnly(ctx)
}

func (c Client) ReadWrite() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ReadWrite(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ReadWrite(ctx)
}

func (c Client) MemoryUsage(key string, samples ...int) *redis.IntCmd {
	if c.useCtx {
		return c.Client.MemoryUsage(c.ctx, key, samples...)
	}
	ctx := context.Background()
	return c.Client.MemoryUsage(ctx, key, samples...)
}

func (c Client) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	if c.useCtx {
		return c.Client.Eval(c.ctx, script, keys, args...)
	}
	ctx := context.Background()
	return c.Client.Eval(ctx, script, keys, args...)
}

func (c Client) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if c.useCtx {
		return c.Client.EvalSha(c.ctx, sha1, keys, args...)
	}
	ctx := context.Background()
	return c.Client.EvalSha(ctx, sha1, keys, args...)
}

func (c Client) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	if c.useCtx {
		return c.Client.ScriptExists(c.ctx, hashes...)
	}
	ctx := context.Background()
	return c.Client.ScriptExists(ctx, hashes...)
}

func (c Client) ScriptFlush() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ScriptFlush(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ScriptFlush(ctx)
}

func (c Client) ScriptKill() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ScriptKill(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ScriptKill(ctx)
}

func (c Client) ScriptLoad(script string) *redis.StringCmd {
	if c.useCtx {
		return c.Client.ScriptLoad(c.ctx, script)
	}
	ctx := context.Background()
	return c.Client.ScriptLoad(ctx, script)
}

func (c Client) Publish(channel string, message interface{}) *redis.IntCmd {
	if c.useCtx {
		return c.Client.Publish(c.ctx, channel, message)
	}
	ctx := context.Background()
	return c.Client.Publish(ctx, channel, message)
}

func (c Client) PubSubChannels(pattern string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.PubSubChannels(c.ctx, pattern)
	}
	ctx := context.Background()
	return c.Client.PubSubChannels(ctx, pattern)
}

func (c Client) PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	if c.useCtx {
		return c.Client.PubSubNumSub(c.ctx, channels...)
	}
	ctx := context.Background()
	return c.Client.PubSubNumSub(ctx, channels...)
}

func (c Client) PubSubNumPat() *redis.IntCmd {
	if c.useCtx {
		return c.Client.PubSubNumPat(c.ctx)
	}
	ctx := context.Background()
	return c.Client.PubSubNumPat(ctx)
}

func (c Client) ClusterSlots() *redis.ClusterSlotsCmd {
	if c.useCtx {
		return c.Client.ClusterSlots(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterSlots(ctx)
}

func (c Client) ClusterNodes() *redis.StringCmd {
	if c.useCtx {
		return c.Client.ClusterNodes(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterNodes(ctx)
}

func (c Client) ClusterMeet(host, port string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterMeet(c.ctx, host, port)
	}
	ctx := context.Background()
	return c.Client.ClusterMeet(ctx, host, port)
}

func (c Client) ClusterForget(nodeID string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterForget(c.ctx, nodeID)
	}
	ctx := context.Background()
	return c.Client.ClusterForget(ctx, nodeID)
}

func (c Client) ClusterReplicate(nodeID string) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterReplicate(c.ctx, nodeID)
	}
	ctx := context.Background()
	return c.Client.ClusterReplicate(ctx, nodeID)
}

func (c Client) ClusterResetSoft() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterResetSoft(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterResetSoft(ctx)
}

func (c Client) ClusterResetHard() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterResetHard(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterResetHard(ctx)
}

func (c Client) ClusterInfo() *redis.StringCmd {
	if c.useCtx {
		return c.Client.ClusterInfo(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterInfo(ctx)
}

func (c Client) ClusterKeySlot(key string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ClusterKeySlot(c.ctx, key)
	}
	ctx := context.Background()
	return c.Client.ClusterKeySlot(ctx, key)
}

func (c Client) ClusterGetKeysInSlot(slot int, count int) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ClusterGetKeysInSlot(c.ctx, slot, count)
	}
	ctx := context.Background()
	return c.Client.ClusterGetKeysInSlot(ctx, slot, count)
}

func (c Client) ClusterCountFailureReports(nodeID string) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ClusterCountFailureReports(c.ctx, nodeID)
	}
	ctx := context.Background()
	return c.Client.ClusterCountFailureReports(ctx, nodeID)
}

func (c Client) ClusterCountKeysInSlot(slot int) *redis.IntCmd {
	if c.useCtx {
		return c.Client.ClusterCountKeysInSlot(c.ctx, slot)
	}
	ctx := context.Background()
	return c.Client.ClusterCountKeysInSlot(ctx, slot)
}

func (c Client) ClusterDelSlots(slots ...int) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterDelSlots(c.ctx, slots...)
	}
	ctx := context.Background()
	return c.Client.ClusterDelSlots(ctx, slots...)
}

func (c Client) ClusterDelSlotsRange(min, max int) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterDelSlotsRange(c.ctx, min, max)
	}
	ctx := context.Background()
	return c.Client.ClusterDelSlotsRange(ctx, min, max)
}

func (c Client) ClusterSaveConfig() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterSaveConfig(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterSaveConfig(ctx)
}

func (c Client) ClusterSlaves(nodeID string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.ClusterSlaves(c.ctx, nodeID)
	}
	ctx := context.Background()
	return c.Client.ClusterSlaves(ctx, nodeID)
}

func (c Client) ClusterFailover() *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterFailover(c.ctx)
	}
	ctx := context.Background()
	return c.Client.ClusterFailover(ctx)
}

func (c Client) ClusterAddSlots(slots ...int) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterAddSlots(c.ctx, slots...)
	}
	ctx := context.Background()
	return c.Client.ClusterAddSlots(ctx, slots...)
}

func (c Client) ClusterAddSlotsRange(min, max int) *redis.StatusCmd {
	if c.useCtx {
		return c.Client.ClusterAddSlotsRange(c.ctx, min, max)
	}
	ctx := context.Background()
	return c.Client.ClusterAddSlotsRange(ctx, min, max)
}

func (c Client) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	if c.useCtx {
		return c.Client.GeoAdd(c.ctx, key, geoLocation...)
	}
	ctx := context.Background()
	return c.Client.GeoAdd(ctx, key, geoLocation...)
}

func (c Client) GeoPos(key string, members ...string) *redis.GeoPosCmd {
	if c.useCtx {
		return c.Client.GeoPos(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.GeoPos(ctx, key, members...)
}

func (c Client) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if c.useCtx {
		return c.Client.GeoRadius(c.ctx, key, longitude, latitude, query)
	}
	ctx := context.Background()
	return c.Client.GeoRadius(ctx, key, longitude, latitude, query)
}

func (c Client) GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	if c.useCtx {
		return c.Client.GeoRadiusStore(c.ctx, key, longitude, latitude, query)
	}
	ctx := context.Background()
	return c.Client.GeoRadiusStore(ctx, key, longitude, latitude, query)
}

func (c Client) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	if c.useCtx {
		return c.Client.GeoRadiusByMember(c.ctx, key, member, query)
	}
	ctx := context.Background()
	return c.Client.GeoRadiusByMember(ctx, key, member, query)
}

func (c Client) GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	if c.useCtx {
		return c.Client.GeoRadiusByMemberStore(c.ctx, key, member, query)
	}
	ctx := context.Background()
	return c.Client.GeoRadiusByMemberStore(ctx, key, member, query)
}

func (c Client) GeoSearch(key string, q *redis.GeoSearchQuery) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.GeoSearch(c.ctx, key, q)
	}
	ctx := context.Background()
	return c.Client.GeoSearch(ctx, key, q)
}

func (c Client) GeoSearchLocation(key string, q *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	if c.useCtx {
		return c.Client.GeoSearchLocation(c.ctx, key, q)
	}
	ctx := context.Background()
	return c.Client.GeoSearchLocation(ctx, key, q)
}

func (c Client) GeoSearchStore(key, store string, q *redis.GeoSearchStoreQuery) *redis.IntCmd {
	if c.useCtx {
		return c.Client.GeoSearchStore(c.ctx, key, store, q)
	}
	ctx := context.Background()
	return c.Client.GeoSearchStore(ctx, key, store, q)
}

func (c Client) GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	if c.useCtx {
		return c.Client.GeoDist(c.ctx, key, member1, member2, unit)
	}
	ctx := context.Background()
	return c.Client.GeoDist(ctx, key, member1, member2, unit)
}

func (c Client) GeoHash(key string, members ...string) *redis.StringSliceCmd {
	if c.useCtx {
		return c.Client.GeoHash(c.ctx, key, members...)
	}
	ctx := context.Background()
	return c.Client.GeoHash(ctx, key, members...)
}

func (c Client) Ctx(ctx context.Context) *Client {
	return &Client{
		Client: c.Client,
		ctx:    ctx,
		useCtx: true,
	}
}

// NewRedisClient return the redis client
func NewRedisClient(conf *Config) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Client{Client: client, useCtx: false}, nil
}
