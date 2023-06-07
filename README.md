# crawler_book
golang爬虫（豆瓣读书）

## Branches
- Branch [simple](https://github.com/Transmigration-zhou/crawler_book/tree/simple) 对应单任务爬虫
- Branch [concurrent](https://github.com/Transmigration-zhou/crawler_book/tree/concurrent) 对应并发版爬虫
- 分布式爬虫正在开发中

## P.S.
1. 使用Elasticsearch进行数据存储（simple分支除外）
```bash
docker run -d -p 9200:9200 elasticsearch
```
2. 有前端展示页面（simple分支除外）

## 架构
### Simple
![Simple](https://github.com/Transmigration-zhou/crawler_book/assets/57855015/c637ff4d-33f4-4dc2-b9be-6a65e366104e)

### Concurrent
![Concurrent](https://github.com/Transmigration-zhou/crawler_book/assets/57855015/e49e891f-e3ab-4382-a614-a79f0b20b6a3)

## Scheduler
### 简单调度器
![SimpleScheduler](https://github.com/Transmigration-zhou/crawler_book/assets/57855015/f4279e58-9fea-4551-b48b-6751b63204f0)

### 并发调度器
![QueuedScheduler](https://github.com/Transmigration-zhou/crawler_book/assets/57855015/7bebdc84-664e-432b-9ae0-6bb95a2645f1)

### 队列实现调度器
![image](https://github.com/Transmigration-zhou/crawler_book/assets/57855015/5cf19e48-dd3d-4507-8f71-ad3b5a179a9f)
