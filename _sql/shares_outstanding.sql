select c.symbol, c.cik, cast(n.value as decimal) shares_outstanding
from companies c
       inner join (select s1.*
                   from subs s1
                          inner join companies c1 on s1.cik = c1.cik
                   where s1.form in ('10-K', '10-Q', '10-K/A', '10-Q/A')
                     and c1.symbol = :symbol
                   order by period desc
                   limit 1) s
       inner join nums n on s.adsh = n.adsh and n.tag = 'EntityCommonStockSharesOutstanding'
where c.symbol = :symbol;


select c.symbol, c.cik, cast(n.value as decimal) shares_outstanding
from companies c
       inner join (select s1.*
                   from subs s1
                   where s1.form in ('10-K', '10-Q', '10-K/A', '10-Q/A')
                   order by period desc
                   limit 1) s
       inner join nums n on s.adsh = n.adsh and n.tag = 'EntityCommonStockSharesOutstanding';


select c.symbol, c.name, c.exchange, c.cik, cast(n.value as decimal) shares_outstanding
from companies c
       inner join (select c.symbol, s.cik, max(s.period) period, s.adsh
                   from subs s
                          inner join companies c on s.cik = c.cik
                   where s.form in ('10-K', '10-Q', '10-K/A', '10-Q/A')
                   group by s.cik) s on c.cik = s.cik
       join nums n on s.adsh = n.adsh and n.tag = 'EntityCommonStockSharesOutstanding'
order by c.symbol;


-- select c.symbol, c.name, c.exchange, c.cik, s.adsh, cast(n.value as decimal) value, n.tag
-- from companies c
--        inner join (select c.symbol, s.cik, max(s.period) period, s.adsh
--                    from subs s
--                           inner join companies c on s.cik = c.cik
--                    where s.form in ('10-K', '10-Q', '10-K/A', '10-Q/A')
--                    group by s.cik) s on c.cik = s.cik
--        join nums n on s.adsh = n.adsh and n.tag = 'EarningsPerShareDiluted' and n.qtrs = '1' and n.ddate = s.period
-- order by c.symbol;

create table latest_sub as
select c.symbol, s.cik, max(s.period) period, s.adsh
from subs s
       inner join companies c on s.cik = c.cik
where s.form in ('10-K', '10-Q', '10-K/A', '10-Q/A')
group by s.cik;

create unique index if not exists idx_ls_cik on latest_sub (cik);

create table if not exists summary
(
  symbol                    text not null,
  name                      text not null,
  exchange                  text not null,
  office                    text,
  industry_title            text,
  cik                       text not null,
  adsh                      text not null,
  eps_basic                 real not null,
  eps_diluted               real not null,
  common_shares_outstanding real not null
);

select c.symbol, c.cik, ls.adsh, cast(n.value as decimal) value
from companies c
       inner join latest_sub ls on c.cik = ls.cik
       inner join nums n on ls.adsh = n.adsh and n.tag = 'EntityCommonStockSharesOutstanding' and n.ddate = ls.period
order by c.symbol;


select c.cik, ls.adsh, cast(n.value as decimal) value
from companies c
       inner join latest_sub ls on c.cik = ls.cik
       join nums n on ls.adsh = n.adsh and n.tag = 'EarningsPerShareBasic' and n.qtrs = '1' and n.ddate = ls.period
order by c.symbol;

select c.cik, ls.adsh, cast(n.value as decimal) value
from companies c
       inner join latest_sub ls on c.cik = ls.cik
       join nums n on ls.adsh = n.adsh and n.tag = 'EarningsPerShareDiluted' and n.qtrs = '1' and n.ddate = ls.period
order by c.symbol;



select *
from nums n
where n.adsh = '0000320193-22-000070'
  and n.tag = 'EarningsPerShareDiluted'
  and n.qtrs = '1';

select s.*
from subs s
where s.form in ('10-K', '10-Q', '10-K/A', '10-Q/A')
  and s.cik = '1342936'
order by period desc;


select *
from subs s
where s.cik = '320193';

select *
from companies c
where c.symbol = 'TSLA';


select *
from subs s
where s.cik = '1318605';

select c.symbol, c.name, s.*
from subs s
       inner join companies c on s.cik = c.cik
where s.form = '10-K/A'

select *
from tags t
where t.tlabel like '%Earnings%'