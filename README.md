# cardcaptor

Blizzard API를 이용하여, Hearthstone 카드 정보를 수집하고 비정규화하여 AWS DynamoDB에 적재한다.

## How to use

```bash
cardcaptor crawl -key <accessKey> [-db <dbpath>]
cardcaptor struct -akid <AWS IAM key> -secret <AWS IAM secret>
```

### Example

```bash
cardcaptor -db="./cards.db" -key="USNtEMfE48HPRJDeX4a0o9PqjhdQfM6TgcA"
```

## Sub-Commands

### crawl

[블리자드 REST API](https://develop.battle.net)로부터 데이터를 수집한다.

- db: 확장자를 포함하는 새로 생성될 db 파일의 경로
- key : [블리자드 API](https://develop.battle.net/access)에서 발급받은 accessKey

### struct

데이터 저장을 위한 AWS DynampDB 테이블을 생성한다.

- akid: (필수) AWS IAM accessID [발급방법](https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/SettingUp.DynamoWebService.html#SettingUp.DynamoWebService.GetCredentials)
- secret: (필수) AWS IAM secretKey
- region: AWS region [region 목록](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html)
- table: 생성할 테이블 이름
- rcu: 프로비전된 읽기 유닛 개수 설정
- wcu: 프로비전된 쓰기 유닛 개수 설정

### store

수집된 데이터로부터 비정규화된 json 포맷을 작성하고 DynamoDB에 저장한다.

- akid: (필수) AWS IAM accessID [발급방법](https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/SettingUp.DynamoWebService.html#SettingUp.DynamoWebService.GetCredentials)
- secret: (필수) AWS IAM secretKey
- region: AWS region [region 목록](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html)
- table: 저장할 테이블 이름
- db: crawl 커맨드를 통해 생성된 sqlite 데이터 파일

## ERD

![hearthstone](https://user-images.githubusercontent.com/43606714/80727008-213c7700-8b40-11ea-9d1c-6acedacad873.png)
