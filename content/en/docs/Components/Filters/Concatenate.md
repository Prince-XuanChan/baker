---
title: "Concatenate"
weight: 10
date: 2020-12-03
---
## Filter *Concatenate*

### Overview
Concatenate up to 10 fields' values to a single field

### Configuration

Keys available in the `[filter.config]` section:

|Name|Type|Default|Required|Description|
|----|:--:|:-----:|:------:|-----------|
| Fields| array of strings| []| false| The field names to concatenate, in order|
| Target| string| ""| false| The field name to save the concatenated value to|
| Separator| string| ""| false| Separator to concatenate the values. Must either be empty or a single ASCII, non-nil char|
