select c.cik,
       c.symbol,
       s.adsh,
       s.fye,
       s.form,
       s.period,
       s.fy,
       s.fp,
       s.detail,
       s.instance
from companies c
       inner join subs s on c.cik = s.cik and s.form in ('10-Q')
where c.symbol = 'AAPL';

select * from (
select n.adsh                   accession_number,
       submission.instance,
       n.qtrs,
       n.ddate                  period_end_date,
       n.tag,
       p.plabel                 preferred_label,
       t.tlabel   taxonomy_label,
       t.datatype               data_type,
       n.uom                    unit_of_measure,
       CAST(n.value as decimal) value,
       p.report,
       p.line
from nums n
       inner join (select s.adsh, s.period, s.instance
                   from companies c
                          inner join subs s on c.cik = s.cik
                   where c.symbol = 'TSLA') submission on n.adsh = submission.adsh and n.ddate = submission.period
       inner join tags t on n.tag = t.tag and n.version = t.version
       inner join pres p on n.adsh = p.adsh and n.tag = p.tag and n.version = p.version
order by p.report, p.line) p where p.report = 4;



select n.*,
       p.*,
       t.*
from nums n
       inner join (select s.adsh, s.period
                   from companies c
                          inner join subs s on c.cik = s.cik and s.form in ('10-Q')
                   where c.symbol = 'AAPL') submission on n.adsh = submission.adsh and n.ddate = submission.period
       inner join tags t on n.tag = t.tag and n.version = t.version
       inner join pres p on n.adsh = p.adsh and n.tag = p.tag and n.version = p.version
order by p.report, p.line;

select s.*, n.*
from nums n
       inner join subs s on n.adsh = s.adsh and s.form = '10-Q'
where n.adsh = '0000320193-22-000070'
  and n.tag = 'EarningsPerShareDiluted'
  and n.qtrs = '1'
order by n.ddate desc
limit 1;


select s.*, n.*
from nums n
       inner join subs s on n.adsh = s.adsh and s.form = '10-K'
where n.adsh = '0001437749-22-021365'



select *
from subs s
where s.form = '10-K';