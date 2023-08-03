# ARP Supression in VXLAN Domains

VXLAN (Virtual Extensible LAN) is a network virtualization technology that enables the creation of virtual Layer 2 networks over an underlying Layer 3 network. VXLAN is commonly used in data centers and cloud environments to provide network segmentation, multi-tenancy, and scalable network designs.

## LAF (Learning and Forwarding) Switches

In the context of VXLAN, LAF (Learning and Forwarding) switches refer to network switches that are VXLAN-aware and have the ability to learn MAC (Media Access Control) addresses from VXLAN encapsulated packets and perform forwarding based on this learned information. LAF switches are integral to the VXLAN overlay network, enabling the proper delivery of traffic between VXLAN endpoints.

##  ARP Suppression in VXLAN:
ARP (Address Resolution Protocol) is used to resolve IP addresses to MAC addresses on a local network. In VXLAN environments, where the Layer 2 segments are extended over a Layer 3 network, ARP packets can be problematic due to the flooding behavior in traditional Ethernet networks. ARP packets can be sent to all hosts in the same VXLAN segment, causing unnecessary broadcast traffic, especially in large-scale deployments.

ARP suppression is a technique employed in VXLAN environments to reduce ARP flooding and limit ARP packets to only the necessary endpoints. The goal is to minimize the broadcast domain for ARP traffic within the VXLAN overlay network, improving network efficiency and reducing unnecessary bandwidth consumption.

## How ARP Suppression Works in VXLAN

When ARP suppression is enabled on LAF switches within the VXLAN overlay network, the LAF switches intercept ARP packets sent by VXLAN endpoints. The LAF switches then perform ARP suppression, allowing them to learn the necessary MAC-to-IP mappings and associate them with specific VXLAN tunnel endpoints.

When an ARP request is received by a LAF switch, it checks its ARP table to see if it already knows the mapping for the requested IP address. If the mapping is found, the switch does not flood the ARP request to all hosts in the VXLAN segment; instead, it responds directly to the requester with the MAC address associated with the requested IP address. This way, ARP traffic is significantly reduced, as ARP requests are limited to the necessary endpoints.

ARP suppression helps improve network performance and reduces unnecessary broadcast traffic within the VXLAN overlay, making VXLAN deployments more scalable and efficient.

Note: The specific implementation of ARP suppression may vary depending on the vendor's equipment and software used in the VXLAN deployment. Different network devices and VXLAN controllers may have their own mechanisms to achieve ARP suppression.
