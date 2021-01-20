# rwxArch Hacking VM Project

Remember this is *specifically* for setting up Arch to run *from within a virtual machine for hacking (not development)* so things like time and boot manager setup are irrelevant.

1. Setup VMware or VirtualBox.
1. Download the vanilla Arch ISO (not `bootstrap`).
1. Validate the ISO sig and checksums.
1. Create and configure the VM.
1. Boot into the *live boot*.
1. Set the `root` password.
1. Update package repo list `pacman -Sy`.
1. Install SSHD `pacman -S openssh` (not `openssh-server`).
1. Enable SSHD `systemctl enable sshd` (not `ssh`).
1. Active SSHD `systemctl start sshd`.
1. Check your IP `ip -br a` then ping from host.
1. Connect from host OS with SSH (using new password).
1. Enable password-less login with `ssh-copy-id <ip>`.
1. Take a snapshot for use later.
1. Create single bootable root partition: `cfdisk /dev/sda` (no GPT)
1. Format and create a filesystem: `mkfs.ext4 /dev/sda1`
1. Mount the partition: `mount /dev/sda1 /mnt`
1. Validate and view partition: `lsblk; df -h`
1. Install essential package with `pacstrap /mnt base linux linux-firmware vim less openssh inetutils netctl which dhcpcd`
1. Generate the `fstab`: `genfstab -P /mnt >> /mnt/etc/fstab`

Do these steps again (forget to arch-chroot last time in new window):

1. Chroot into the new fs: `arch-chroot /mnt`
1. Setup timezone: `ln -sf /usr/share/zoneinfo/Region/City /etc/localtime`
1. Check timezone `timedatectl status`
1. Fix the time: `timedatectl set-ntp true`
1. Check NTP servers: `timedatectl show-timesync --all`
1. Check the hardware time clock: `hwclock`
1. Setup locale: `locale-gen`


