select cf.id,
    cf.caption,
    cf.[type],
    cf.owner,
    cf.tab,
    cf.[unique],
    cf.mandatory,
    cf.[order],
    isnull(cf.extra, '') extra,
    isnull(cfv.value, '') value
from dbo.CustomFields cf
left join dbo.CustomFieldsValues cfv on cfv.field_id = cf.id and cf.owner = 'Person'
where cfv.owner_id = ?
order by cf.[order]
