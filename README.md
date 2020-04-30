# cardcaptor

Blizzard API를 이용하여, Hearthstone 카드 정보를 수집하고 SQLite3(\*.db)에 적재한다.

## How to use

```bash
cardcaptor -db=<dbpath> -key=<accessKey>
```

### Example

```bash
cardcaptor -db="./cards.db" -key="USNtEMfE48HPRJDeX4a0o9PqjhdQfM6TgcA"
```

### Arguments

- db: 확장자를 포함하는 새로 생성될 db 파일의 경로
- key : [블리자드 API](https://develop.battle.net/access)에서 발급받은 accessKey
