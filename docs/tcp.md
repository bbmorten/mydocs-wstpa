# TCP Protocol

## TCP Utilities

### [tcpflow](https://github.com/simsong/tcpflow)

Can be installed with `brew install tcpflow`

###  Every Linux networking tool I know

![Every Linux networking tool I know](FyZhNHsWcAE_FsU.jpeg)

##  MacOS Parameters for TCP

### net.inet.tcp.mssdflt: 512

This parameter specifies the default maximum segment size (MSS) for TCP connections. The MSS is the largest amount of data that can be transmitted in a single TCP segment. This value is used by TCP when negotiating the MSS during the TCP handshake.

### net.inet.tcp.keepidle: 7200000

This parameter specifies the amount of time that a TCP connection can remain idle before TCP sends a keepalive probe. A keepalive probe is a message that is sent by TCP to check if the remote endpoint is still reachable. If a response is not received after a certain number of probes, the connection is considered dead.

### net.inet.tcp.keepintvl: 75000

This parameter specifies the interval between keepalive probes sent by TCP. The default value is 75 seconds.

### net.inet.tcp.sendspace: 131072

This parameter specifies the size of the send buffer used by TCP for each socket. The send buffer is a data structure that stores data that is waiting to be sent. Increasing the send buffer size can help improve network performance for applications that generate a lot of network traffic.

### net.inet.tcp.recvspace: 131072

This parameter specifies the size of the receive buffer used by TCP for each socket. The receive buffer is a data structure that stores data that has been received but not yet read by the application. Increasing the receive buffer size can help improve network performance for applications that receive a lot of network traffic.

### net.inet.tcp.keepinit: 75000

This parameter specifies the time interval between the transmission of the first and second keepalive probes on a TCP connection. The default value is 75 seconds.

### net.inet.tcp.v6mssdflt: 1024

This parameter specifies the default maximum segment size (MSS) for TCP connections over IPv6. The MSS is the largest amount of data that can be transmitted in a single TCP segment. This value is used by TCP when negotiating the MSS during the TCP handshake.

### net.inet.tcp.reass.overflows: 0

This parameter indicates the number of times the TCP reassembly queue has overflowed due to too many out-of-order packets. If this value is non-zero, it may indicate that the system is under heavy network load or that the TCP reassembly code needs to be tuned.

### net.inet.tcp.reass.qlen: 0

This parameter specifies the maximum number of TCP segments that can be queued for reassembly at any one time. If this value is non-zero, it may indicate that the system is under heavy network load or that the TCP reassembly code needs to be tuned.

### net.inet.tcp.log.privacy: 1

This parameter enables or disables logging of IP addresses in TCP connection logs. If this value is set to 1, IP addresses will be anonymized in the log files.

### net.inet.tcp.log.rate_limit: 1000

This parameter specifies the maximum number of log messages that can be generated per second by TCP. If the rate of log messages exceeds this limit, TCP will start dropping log messages.

### net.inet.tcp.log.rate_duration: 60

This parameter specifies the time interval in seconds over which the TCP log rate limit is enforced. If the rate of log messages exceeds the limit during this interval, TCP will start dropping log messages.

### net.inet.tcp.log.rate_max: 0

This parameter specifies the maximum number of log messages that TCP will allow per time interval. If this value is set to 0, there is no maximum limit.

### net.inet.tcp.log.rate_exceeded_total: 0

This parameter specifies the total number of log messages that TCP has dropped due to exceeding the log rate limit.

### net.inet.tcp.log.rate_current: 0

This parameter specifies the current log message rate. If the rate exceeds the log rate limit, messages will be dropped.

### net.inet.tcp.log.enable: 0

This parameter enables or disables TCP connection logging.

### net.inet.tcp.log.enable_usage: connection:0x00000001 rtt:0x00000002 ka:0x00000004 log:0x00000008 loop:0x00000010 local:0x00000020 gw:0x00000040 syn:0x00000100 fin:0x00000200 rst:0x00000400 dropnecp:0x00001000 droppcb:0x00002000 droppkt:0x00004000 fswflow:0x00008000 state:0x00010000 synrxmt:0x00020000 output:0x00040000

This parameter specifies the type of TCP connection information that will be logged. Each bit in the value corresponds to a different type of information, such as RTT measurements, keepalive probes, or connection state changes.

### net.inet.tcp.log.rtt_port: 0

This parameter specifies the port number used for TCP RTT measurements. If set to 0, the system will use a default port number.

### net.inet.tcp.log.thflags_if_family: 0

This parameter enables or disables logging of TCP packet header flags (such as SYN or FIN) based on the interface family.

### net.inet.tcp.cubic_tcp_friendliness: 0

This parameter enables or disables TCP friendliness in the CUBIC congestion control algorithm. TCP friendliness ensures that CUBIC does not monopolize network resources and allows other TCP connections to share the available bandwidth.

### net.inet.tcp.cubic_fast_convergence: 0

This parameter enables or disables fast convergence in the CUBIC congestion control algorithm. Fast convergence allows CUBIC to quickly adapt to changes in network conditions.

### net.inet.tcp.cubic_use_minrtt: 0

This parameter enables or disables the use of the minimum RTT in the CUBIC congestion control algorithm. Using the minimum RTT can help improve performance over high-latency networks.

### net.inet.tcp.cubic_minor_fixes: 1

This parameter enables or disables minor fixes in the CUBIC congestion control algorithm. These fixes are intended to improve the algorithm's performance.

### net.inet.tcp.cubic_rfc_compliant: 1

This parameter enables or disables compliance with the RFC for the CUBIC congestion control algorithm.

### net.inet.tcp.bg_target_qdelay: 40

This parameter specifies the target queue delay for background traffic in the TCP Background Transport (TBT) congestion control algorithm.

### net.inet.tcp.bg_allowed_increase: 8

This parameter specifies the maximum allowed increase in congestion window size for background traffic in the TCP Background Transport (TBT) congestion control algorithm.

### net.inet.tcp.bg_tether_shift: 1

This parameter specifies the number of bits to shift the congestion window for background traffic in the TCP Background Transport (TBT) congestion control algorithm.

### net.inet.tcp.bg_ss_fltsz: 2

This parameter specifies the size of the slow start filter in the TCP Background Transport (TBT) congestion control algorithm.

### net.inet.tcp.ledbat_plus_plus: 1

This parameter enables or disables the LEDBAT congestion control algorithm. LEDBAT is designed to reduce network congestion and improve

### net.inet.tcp.rledbat

Determines if RLEDBAT congestion control algorithm is enabled. RLEDBAT stands for Rate Limited Enhanced Delay-Based Algorithm for TCP, which is a congestion control algorithm that aims to avoid network congestion and reduce latency.

### net.inet.tcp.cc_debug

Determines whether TCP congestion control debug information is logged. When set to 1, debug information about the congestion control is logged to the kernel log.

### net.inet.tcp.newreno_sockets

Determines whether NewReno congestion control algorithm is enabled for TCP sockets. NewReno is a congestion control algorithm that is designed to overcome problems with the original TCP Reno algorithm.

### net.inet.tcp.background_sockets

Specifies the number of TCP sockets that are considered "background" sockets. These sockets are subject to different congestion control algorithms and are not given priority over other sockets.

### net.inet.tcp.use_ledbat

Determines if LEDBAT congestion control algorithm is enabled. LEDBAT stands for Low Extra Delay Background Transport, which is a congestion control algorithm designed for low-priority, background traffic.

### net.inet.tcp.cubic_sockets

Specifies the number of TCP sockets that use the CUBIC congestion control algorithm. CUBIC is a congestion control algorithm that aims to improve performance over high-speed, long-distance networks.

### net.inet.tcp.use_newreno

Determines if the NewReno congestion control algorithm is enabled for all TCP sockets.

### net.inet.tcp.mptcp_preferred_version

Specifies the preferred version of the Multipath TCP (MPTCP) protocol. MPTCP is a TCP extension that allows a single TCP connection to use multiple paths simultaneously, improving performance and reliability.

### net.inet.tcp.backoff_maximum

Specifies the maximum number of milliseconds that TCP will back off after a retransmission timeout. The default value is 64 seconds.

### net.inet.tcp.ecn_timeout

Specifies the number of seconds that ECN (Explicit Congestion Notification) marking will be enabled for a TCP connection.

### net.inet.tcp.disable_tcp_heuristics

Disables TCP heuristics. When TCP heuristics are enabled, the operating system tries to determine the optimal TCP parameters based on the characteristics of the network.

### net.inet.tcp.mptcp_version_timeout

Specifies the number of seconds that MPTCP versions are kept in the cache.

### net.inet.tcp.clear_tfocache

Clears the TCP fast open (TFO) cache.

### net.inet.tcp.flow_control_response

Determines whether TCP flow control responses are enabled. When set to 1, TCP flow control responses are sent to the remote host to indicate when the local host is congested.

### net.inet.tcp.log_in_vain

Specifies whether TCP logging is enabled for packets that arrive at the host for which there is no listening TCP process.

### net.inet.tcp.ack_strategy

Specifies the TCP acknowledgment (ACK) strategy. The default value of 1 specifies delayed ACKs, where a single ACK is sent for every other data packet received.

### net.inet.tcp.blackhole

Specifies whether the operating system will drop packets that are not part of a known TCP connection. When set to 1, packets that do not belong to a known connection are dropped.

### net.inet.tcp.aggressive_rcvwnd_inc

Specifies whether TCP will aggressively increase the receive window size. When set to 1, the receive window size is increased more quickly.

### net.inet.tcp.delayed_ack

Specifies the maximum time in milliseconds that TCP will delay sending an ACK packet.

### net.inet.tcp.recvbg

Specifies whether TCP background receiver processing is enabled. When set to 1, the operating system performs background processing on TCP receiver sockets.

### net.inet.tcp.recv_throttle_minwin

The `net.inet.tcp.recv_throttle_minwin` parameter sets the minimum receive window size in bytes that a TCP connection should use. It is used to prevent TCP connections from overwhelming a receiver that is unable to keep up with the incoming data.

### net.inet.tcp.enable_tlp

The `net.inet.tcp.enable_tlp` parameter enables or disables TCP Time-Loss Protection (TLP), which is a mechanism that helps detect and recover from packet loss more quickly. When TLP is enabled, TCP will send a small number of duplicate packets when it suspects that a packet may have been lost, allowing the sender to receive an ACK more quickly and reducing the amount of time spent waiting for a retransmission.

### net.inet.tcp.sack

The `net.inet.tcp.sack` parameter enables or disables the Selective Acknowledgment (SACK) option in TCP. SACK allows a receiver to inform the sender of which packets have been received successfully and which have been lost, allowing the sender to retransmit only the lost packets instead of retransmitting the entire data stream.

### net.inet.tcp.sack_maxholes

The `net.inet.tcp.sack_maxholes` parameter sets the maximum number of SACK holes that can be stored for a TCP connection. A SACK hole is a range of packets that have been lost and need to be retransmitted. When the number of SACK holes exceeds this limit, older holes are discarded.

### net.inet.tcp.sack_globalmaxholes

The `net.inet.tcp.sack_globalmaxholes` parameter sets the maximum number of SACK holes that can be stored across all TCP connections on the system. When the number of SACK holes exceeds this limit, older holes are discarded.

### net.inet.tcp.fastopen_key

The `net.inet.tcp.fastopen_key` parameter is used to enable TCP Fast Open, which is a feature that allows a client to send data in the initial SYN packet of a TCP connection. This can reduce the latency of short-lived connections, such as those used for web browsing. The `fastopen_key` is a secret key that is used to protect against spoofed packets.

### net.inet.tcp.fastopen_backlog

The `net.inet.tcp.fastopen_backlog` parameter sets the size of the queue that holds Fast Open requests waiting for a SYN-ACK from the server.

### net.inet.tcp.fastopen

The `net.inet.tcp.fastopen` parameter enables or disables TCP Fast Open. When set to 0, Fast Open is disabled. When set to 1, Fast Open is enabled for clients only. When set to 2, Fast Open is enabled for both clients and servers.

### net.inet.tcp.now_init

The `net.inet.tcp.now_init` parameter sets the initial value of the TCP timestamp clock. This value is used to generate timestamps for TCP packets and is typically set to the current system time.

### net.inet.tcp.microuptime_init

The `net.inet.tcp.microuptime_init` parameter sets the initial value of the microsecond uptime clock. This clock is used to measure the amount of time that has elapsed since the system was last booted.

### net.inet.tcp.minmss

The `net.inet.tcp.minmss` parameter sets the minimum segment size that TCP will advertise to its peer. This parameter is used to avoid sending packets that are too small and can cause excessive overhead.

### net.inet.tcp.pcbcount

The `net.inet.tcp.pcbcount` parameter shows the current number of active TCP connections on the system.

### net.inet.tcp.tw_pcbcount: 4

This parameter specifies the maximum number of TCP connections that are allowed to remain in the TIME_WAIT state. Once a TCP connection is closed, it enters the TIME_WAIT state for a period of time to ensure that any delayed packets related to the closed connection are not misinterpreted as belonging to a new connection.

### net.inet.tcp.icmp_may_rst: 1

This parameter specifies whether TCP may use ICMP to reset connections. ICMP is an Internet protocol used by routers and other devices to send error messages indicating that a requested service is not available or that a packet has been dropped.

### net.inet.tcp.rtt_min: 100

This parameter specifies the minimum round-trip time (RTT) value in milliseconds for TCP connections. This value is used as a baseline for calculating retransmission timeouts.

### net.inet.tcp.rexmt_slop: 200

This parameter specifies the amount of extra time, in milliseconds, that should be added to the calculated retransmission timeout value. This is to ensure that retransmission timeouts are not triggered too early due to transient network congestion or delays.

### net.inet.tcp.randomize_ports: 0

This parameter specifies whether TCP should randomize the port numbers used for outgoing connections. When set to 1, the system will use a random high-numbered port for each outgoing connection to make it harder for attackers to guess the port number being used.

### net.inet.tcp.win_scale_factor: 3

This parameter specifies the window scale factor for TCP connections. The window scale factor is used to increase the maximum window size that can be used for a connection, allowing for larger amounts of data to be transmitted without requiring frequent acknowledgments.

### net.inet.tcp.init_rtt_from_cache: 1

This parameter specifies whether TCP should initialize the initial round-trip time (RTT) estimate from the cache. The cache stores information about previous connections and can help speed up the process of establishing new connections.

### net.inet.tcp.tso_debug: 0

This parameter specifies whether TCP segmentation offload (TSO) debugging is enabled. TSO is a feature that allows TCP to offload the segmentation of data to the network interface card (NIC), which can improve performance by reducing CPU overhead.

### net.inet.tcp.rxt_seg_max: 1024

This parameter specifies the maximum number of retransmitted segments that can be outstanding for a TCP connection. If this value is exceeded, the connection will be reset.

### net.inet.tcp.rxt_seg_drop: 0

This parameter specifies whether TCP should drop retransmitted segments that have already been sent. This can help reduce network congestion and improve performance.

### net.inet.tcp.tcbhashsize: 4096

This parameter specifies the size of the hash table used to store TCP control blocks (TCBs), which represent active TCP connections. A larger hash table can improve performance by reducing the time required to locate a TCB.

### net.inet.tcp.keepcnt: 8

This parameter specifies the maximum number of keepalive probes that can be sent for a TCP connection that is idle. Keepalive probes are packets that are sent periodically to ensure that a connection remains active.

### net.inet.tcp.msl: 15000

This parameter specifies the maximum segment lifetime (MSL) value for TCP connections. The MSL is the time period after which a TCP segment is assumed to have been lost if no acknowledgement has been received.

### net.inet.tcp.fin_timeout: 60000

This parameter specifies the time period, in milliseconds, that a TCP connection will remain in the FIN_WAIT_2 state before being closed. This state is entered when a connection has been

### net.inet.tcp.fin_timeout

This parameter sets the amount of time that a connection in the "FIN-WAIT-2" state remains open if it is not closed by the remote endpoint. The default value is 60 seconds.

### net.inet.tcp.max_persist_timeout

This parameter sets the maximum amount of time that TCP will continue to send zero-window probes to a peer in order to attempt to clear a zero window. If this value is set to zero, TCP will use a default value of 5 minutes.

### net.inet.tcp.always_keepalive

This parameter controls whether TCP will send keepalive messages on a connection even if the connection is idle. If this parameter is set to 1, keepalive messages will always be sent. The default value is 0.

### net.inet.tcp.timer_fastmode_idlemax

This parameter controls the maximum amount of idle time that a TCP connection can have before TCP switches to fast timer mode. In fast timer mode, TCP timers fire more frequently in order to speed up retransmission of lost packets. The default value is 10 seconds.

### net.inet.tcp.broken_peer_syn_rexmit_thres

This parameter sets the number of times that TCP will retransmit a SYN packet before assuming that the peer is broken and giving up on the connection. The default value is 5.

### net.inet.tcp.tcp_timer_advanced

This parameter controls the number of times that the TCP timer thread will advance the timer list in a single call. This is done in order to reduce the overhead of timer processing. The default value is 247.

### net.inet.tcp.tcp_resched_timerlist

This parameter controls the number of times that the TCP timer thread will reschedule the timer list in a single call. This is done in order to reduce the overhead of timer processing. The default value is 122571.

### net.inet.tcp.pmtud_blackhole_detection

This parameter enables or disables detection of black hole paths in Path MTU Discovery (PMTUD). If this parameter is set to 1, TCP will attempt to detect black hole paths and reduce the PMTU accordingly. The default value is 1.

### net.inet.tcp.pmtud_blackhole_mss

This parameter sets the maximum segment size (MSS) that TCP will use when detecting black hole paths in PMTUD. If the MSS is set too high, it can lead to larger packets being fragmented and potentially dropped by the network. The default value is 1200 bytes
