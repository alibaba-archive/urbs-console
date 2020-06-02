import moment from 'moment';

export function formatTableTime(time: string) {
  const momentT = time ? moment(time) : moment.unix(0);
  return momentT.format('YYYY/MM/DD ah:mm:ss');
}

export function formatTimestamp(time: number) {
  const momentT = moment.unix(time);
  return momentT.format('YYYY/MM/DD ah:mm:ss');
}