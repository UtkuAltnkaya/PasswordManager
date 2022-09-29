import Data from 'models/data';
import User from 'models/user';
import { useMemo } from 'react';
import AddItem from './AddItem';
import style from './Main.module.scss';
import MainItem from './MainItem';

type Props = {
  data: Data<User>;
};

const colors = {
  red: '#EB1D36',
  purple: '#645CAA',
  navyBlue: '#25316D',
  green: '#647E68',
  orange: '#D07000',
  darkMauve: '#874C62',
  brown: '#A77979',
  random: function () {
    const keys = Object.keys(this);
    return this[keys[((keys.length - 1) * Math.random()) << 0]];
  },
};
const color = colors.random();

const Main = ({ data }: Props) => {
  const items = useMemo(() => {
    return data.message.password_list.sort((a: any, b: any) =>
      a.password_name.name < b.password_name.name ? -1 : a.password_name.name > b.password_name.name ? 1 : 0,
    );
    // eslint-disable-next-line
  }, [data.message.password_list.length]);

  return (
    <div className={style.main}>
      <AddItem color={color} />
      {items.map((item) => (
        <MainItem key={item.id} pass={item} color={color} />
      ))}
    </div>
  );
};

export default Main;
