
```md

ClickHouse 默认使用以下端口：
1. TCP 端口（客户端连接）：  
9000
用于 ClickHouse 客户端、驱动程序（如 JDBC、Python 等）连接。
2. HTTP 端口（REST API 和 Web UI）：  
8123
用于通过 HTTP 协议访问 ClickHouse，例如使用 curl 或浏览器访问。
3. HTTPS 端口（可选，启用 SSL 后）：  
8443
如果启用了 HTTPS 支持，则使用此端口。
4. Native secure TCP（启用 TLS 后）：  
9440
启用加密的 TCP 连接时使用。
5. Interserver HTTP 端口（节点间通信）：  
9009
用于分布式查询时节点之间的数据传输。

你可以在配置文件中修改这些端口，配置文件路径通常为：
/etc/clickhouse-server/config.xml
例如，修改 TCP 端口配置项：
<tcp_port>9000</tcp_port>

建议在生产环境中根据需要调整端口并配置防火墙规则以增强安全性。

```