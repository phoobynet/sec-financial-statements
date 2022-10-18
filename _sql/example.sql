select c.*, s.adsh accession_number, si.office, si.industry_title
from companies c
       inner join subs s on c.cik = s.cik
       left join sics si on s.sic = si.code
where c.symbol = 'AAPL';

-- search query
select c.cik, c.symbol, c.name, c.exchange, si.office, si.industry_title industry
from companies c
       inner join subs s on c.cik = s.cik
       left join sics si on s.sic = si.code
where c.symbol = :query
   or c.name like '%Appl%'