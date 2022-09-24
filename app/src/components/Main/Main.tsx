import Data from 'models/data';
import User from 'models/user';
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
  // const color = useMemo(() => colors.random(), [colors]);

  return (
    <>
      <div className={style.main}>
        <AddItem color={color} />
        <div className={style.itemDiv}>
          {data.message.password_list.map((item) => (
            <MainItem key={item.id} pass={item} color={color} />
          ))}
        </div>
      </div>
    </>
  );
};

export default Main;
