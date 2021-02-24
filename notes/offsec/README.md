# Offensive Security Notes

## Self-Training Strategy

First a few core tenets that influence my decisions:

* Build upon core skills unrelated to offsec.
  * Know how to *really* use you computer
  * 

## Seven Steps of Pentesting

1. Information Gathering
2. Reconnaissance
3. Discovery and Scanning
4. Vulnerability Assessment
5. Exploitation
6. Final Analysis and Review
7. Utilize the Testing Results

## Discover and Scanning

Some stupid operating systems (including Ubuntu) depend on the following
`snap` configuration to even get `nmap` to work:

```
sudo snap connect nmap:network-control
```

