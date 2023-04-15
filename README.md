# gofs
```
  _____  ____  ______ _____    
 / ____|/ __ \|  ____/ ____|   
| |  __| |  | | |__ | (___     
| | |_ | |  | |  __| \___ \    
| |__| | |__| | |    ____) |   
 \_____|\____/|_|   |_____/    
```

### Feature

After further study, I deepened my understanding of the distributed system and prepared to rebuild the project.

### Introduce

Before that, let's learn some concepts.

**MetaServer**: It is used to store the metadata of the file system and schedule the system, including directory trees, configurations, logs, etc.

**ChunkServer**: It is used to store the actual file data, usually a fragment of the file.

**gofsctl**: It is used to operate the file system and update configuration and other functions.

As you can see, the whole project is divided into three parts. Let's get more.

**Master&Slave**: MetaServer is the most important component, which determines whether the system can work properly, so we need to achieve high availability on MetaServer and adopt master-slave architecture.

**Namespace**: We hope that there will be diversity in the file system, such as different configurations, file isolation, and each namespace is isolated from each other.

**EditLog**: Each operation of gofs can be regarded as a edit log. In the master-slave architecture, master-slave consistency is achieved by copying and applying EditLog.

**Snapshot**: Use the snapshot to quickly restore the state of gofs to the state before the shutdown. The snapshot will record the current EditLog number and minimize metadata loss with EditLog.

**Heartbeat**: MetaServer periodically communicates with each ChunkServer server through heartbeat information, sends instructions to each ChunkServer server, and receives the status information of each ChunkServer server.

**Chunk**: Usually, we divide files into fixed-size chunks, which is conducive to disk or network I/O and management, and disaster recovery can be achieved through backup blocks.



