Implement instrument
- It will have Docker-compose like YAML interface (a bit intermixed with Ansible inventory perhaps for multi server purposes)
- it will raise automatically vagrant objects at target servers
- it will store last applied state at target servers like helm, in order to avoid doing operations which are already done
- it will automatically delete objects no longer present in state, by quering vagrant API