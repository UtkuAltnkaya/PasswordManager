import { useRef, useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { toast } from 'react-toastify';
import style from './Main.module.scss';

import { RequestToServer } from 'services/Api';
import { Settings } from 'icons/MainIcons';
import useOnClickOutside from 'hooks/useOnClickOutside';

import Data from 'models/data';
import Password from 'models/password';
import SettingsComponent from './Settings';

type Props = {
  pass: Password;
  color: string;
};

const MainItem = ({ pass, color }: Props) => {
  const [openPassword, setOpenPassword] = useState(false);
  const [openSettings, setOpenSetting] = useState(false);
  const ref = useRef(null);
  useOnClickOutside(ref, () => setOpenSetting(false));

  const { data: item, refetch } = useQuery<Data<Password>>(
    ['site', pass.password_name.name],
    () => RequestToServer({ url: 'getpassword', params: { site: pass.password_name.name } }),
    {
      enabled: false,
    },
  );

  const handleClick = () => {
    if (!openPassword) {
      setOpenPassword(true);
      refetch();
      return;
    }
    setOpenPassword(false);
  };
  const handleCopy = () => {
    if (openPassword) {
      navigator.clipboard.writeText(item?.message?.password);
      toast.info('Password copied to click board');
    }
  };
  const closeSettings = () => {
    setOpenSetting(false);
    setOpenPassword(false);
  };

  return (
    <div key={pass.id} className={style.item} style={{ backgroundColor: color }} ref={ref}>
      <div className={style.title} style={{ cursor: openPassword ? 'pointer' : 'auto' }} onClick={handleCopy}>
        {pass.password_name.name}
      </div>
      <div className={style.password}>
        <div onClick={handleClick}>{openPassword ? item?.message?.password : '*****************'}</div>
        <div onClick={() => setOpenSetting(!openSettings)}>
          <Settings width={20} height={20} />
        </div>
      </div>

      {openSettings && <SettingsComponent name={pass.password_name.name} closeSettings={closeSettings} />}
    </div>
  );
};

export default MainItem;
