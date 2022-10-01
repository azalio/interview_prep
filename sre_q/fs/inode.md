# Inode

<https://syedali.net/2015/02/08/understanding-inodes/>

<https://tldp.org/LDP/tlk/fs/filesystem.html>

![inode](https://tldp.org/LDP/tlk/fs/ext2_inode.gif)

## Inode content

```bash
Inode Number
Uid
Gid
Size
Atime
Mtime
Ctime
Blocksize
Mode
Number of links
ACLs
```

Filenames are not stored in inodes, instead they are stored in the data portion of a directory.

## Delete file

When you delete a file the unlink() system call removes the directory entry for the inode and marks it available. The data blocks themselves are not deleted.

## Soft and hard links

The number of links to a file is maintained in an inode. Each time a hard link is created the number of links increases. Soft links do not increase the number of links to a file or directory.

## Superblock

Superblock contains metadata about a filesystem. A filesystem typically stores many copies of a superblock in case one of them gets damaged. Some of the information in a superblock is:

– Filesystem size
– Block size
– Empty and filled blocks
– Size and location of inode table
– Disk block map

You can read superblock information using the command ‘dumpe2fs /dev/mount | grep -i superblock’.

## EXT2 Directories

![directory](https://tldp.org/LDP/tlk/fs/ext2_dir.gif)

In the EXT2 file system, directories are special files that are used to create and hold access paths to the files in the file system.

A directory file is a list of directory entries, each one containing the following information:

**inode**

The inode for this directory entry.
This is an index into the array of inodes held in the Inode Table of the Block Group. In figure  9.3, the directory entry for the file called file has a reference to inode number i1,

**name length**
The length of this directory entry in bytes,

**name**
The name of this directory entry.

## Finding a File in an EXT2 File System

A Linux filename has the same format as all Unix TM filenames have. It is a series of directory names separated by forward slashes (``/'') and ending in the file's name.

One example filename would be /home/rusling/.cshrc where /home and /rusling are directory names and the file's name is .cshrc. Like all other Unix TM systems, Linux does not care about the format of the filename itself; it can be any length and consist of any of the printable characters. To find the inode representing this file within an EXT2 file system the system must parse the filename a directory at a time until we get to the file itself.

The first inode we need is the inode for the root of the file system and we find its number in the file system's superblock.

To read an EXT2 inode we must look for it in the inode table of the appropriate Block Group.

If, for example, the root inode number is 42, then we need the 42nd inode from the inode table of Block Group 0. The root inode is for an EXT2 directory, in other words the mode of the root inode describes it as a directory and it's data blocks contain EXT2 directory entries.

home is just one of the many directory entries and this directory entry gives us the number of the inode describing the /home directory.

We have to read this directory (by first reading its inode and then reading the directory entries from the data blocks described by its inode) to find the rusling entry which gives us the number of the inode describing the /home/rusling directory.

Finally we read the directory entries pointed at by the inode describing the /home/rusling directory to find the inode number of the .cshrc file and from this we get the data blocks containing the information in the file.
