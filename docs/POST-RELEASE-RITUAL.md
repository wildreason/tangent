# Post-Release Ritual

## Rules

1. Never force-push tags
2. Never delete published tags
3. Increment version on conflicts (beta.1 → beta.2)
4. Verify proxy after 5min wait

## Checklist

```bash
# 1. Tag
git tag vX.X.X
git push origin vX.X.X

# 2. Wait 5 minutes

# 3. Verify proxy
curl -s "https://proxy.golang.org/github.com/wildreason/tangent/@v/vX.X.X.info"
git log -1 --format="%ai" vX.X.X  # Timestamps must match

# 4. Test consumer
go get github.com/wildreason/tangent@vX.X.X
```

## If Proxy Cache Stale

Increment version:
```bash
git tag vX.X.X+1 <commit>
git push origin vX.X.X+1
```

Consumer emergency fix:
```bash
GOPROXY=direct go get github.com/wildreason/tangent@vX.X.X
```
