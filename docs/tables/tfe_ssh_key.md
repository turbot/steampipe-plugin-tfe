---
title: "Steampipe Table: tfe_ssh_key - Query Terraform Enterprise SSH Keys using SQL"
description: "Allows users to query SSH Keys in Terraform Enterprise, specifically the key details, providing insights into key management and potential security issues."
---

# Table: tfe_ssh_key - Query Terraform Enterprise SSH Keys using SQL

Terraform Enterprise is a collaborative, scalable, and enterprise-friendly service offered by HashiCorp that provides workflow and automation control for infrastructure as code (IaC). It supports the management of SSH keys, which are essential for secure server communication and access control. These keys are crucial for the security of your infrastructure, and Terraform Enterprise provides a way to manage and monitor them.

## Table Usage Guide

The `tfe_ssh_key` table provides insights into SSH keys within Terraform Enterprise. As a DevOps engineer, explore key-specific details through this table, including key ID, name, and associated metadata. Utilize it to uncover information about keys, such as their creation and last usage time, which can be crucial for auditing and security purposes.

## Examples

### Basic info
Explore the SSH keys in your environment to understand their role and usage. This is crucial for maintaining secure access to your systems.

```sql
select
  *
from
  tfe_ssh_key;
```

### Get SSH key by ID
Identify specific SSH keys by their unique identifiers to quickly access and manage them, enhancing security and control over your system.

```sql
select
  *
from
  tfe_ssh_key
where
  id = 'sshkey-1NSDCvowf3WtbStu';
```