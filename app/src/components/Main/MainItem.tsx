import { useQuery } from '@tanstack/react-query';
import Data from 'models/data';
import Password from 'models/password';
import { useState } from 'react';
import { toast } from 'react-toastify';
import { RequestToServer } from 'services/Api';
import style from './Main.module.scss';

type Props = {
  pass: Password;
  color: string;
};

const MainItem = ({ pass, color }: Props) => {
  const [open, setOpen] = useState(false);

  const { data: item, refetch } = useQuery<Data<Password>>(
    ['site', pass.password_name.name],
    () => RequestToServer({ url: 'getpassword', params: { site: pass.password_name.name } }),
    {
      enabled: false,
    },
  );

  const handleClick = () => {
    if (!open) {
      setOpen(true);
      refetch();
      return;
    }
    setOpen(false);
  };
  const handleCopy = () => {
    if (open) {
      navigator.clipboard.writeText(item?.message?.password);
      toast.info('Password copied to click board');
    }
  };

  return (
    <div key={pass.id} className={style.item} style={{ backgroundColor: color }}>
      <div className={style.title} style={{ cursor: open ? 'pointer' : 'auto' }} onClick={handleCopy}>
        {pass.password_name.name}
      </div>
      <div onClick={handleClick} className={style.password}>
        {open ? item?.message?.password : '*****************'}
      </div>
    </div>
  );
};

export default MainItem;
