import moment from 'moment';

export function formatTableTime (time: string) {
  const momentT = time ? moment(time) : moment();
  return momentT.format('YYYY/MM/DD ah:mm:ss');
}