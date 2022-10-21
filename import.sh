#!/bin/zsh

sqlite3 test.db <<EOF
.mode tabs
.import /Volumes/yotta_1tb/sec_data/2022q3/sub.txt subs
.import /Volumes/yotta_1tb/sec_data/2022q3/tag.txt tags
.import /Volumes/yotta_1tb/sec_data/2022q3/pre.txt pre
.import /Volumes/yotta_1tb/sec_data/2022q3/num.txt nums
create index if not exists idx_subs_cik on subs (cik);
create index if not exists idx_subs_adsh on subs (adsh);
create index if not exists idx_nums_adsh on nums (adsh);
create index if not exists idx_pre_adsh on pre (adsh);
EOF
