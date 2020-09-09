import moment from 'moment'

export function utc2local(dateStr) {
  return moment.utc(dateStr).local().format('YYYY-MM-DD HH:mm:ss')
}
