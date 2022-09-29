import { MainContext } from 'context/provider';
import { AddIcon, DeleteIcon } from 'icons/MainIcons';
import { useContext, useState } from 'react';
import style from './Main.module.scss';

type Props = {
  name: string;
  closeSettings: Function;
};

const Settings = ({ name, closeSettings }: Props) => {
  const [size, setSize] = useState(32);
  const { setUpdateModal, setDeleteModal } = useContext(MainContext);

  const openUpdateModal = () => {
    setUpdateModal({ open: true, site: name, size });
    closeSettings();
  };

  const openDeleteModal = () => {
    setDeleteModal({ open: true, site: name });
    closeSettings();
  };

  return (
    <div className={style.mainSettings}>
      <div className={style.arrow}></div>

      <div className={style.titles}>
        <div className={style.add}>Add</div>
        <div className={style.update} onClick={openUpdateModal}>
          Update
        </div>
        <div className={style.delete} onClick={openDeleteModal}>
          Delete
        </div>
      </div>

      <div className={style.items}>
        <div>
          <AddIcon />
        </div>
        <select value={size} onChange={(e) => setSize(parseInt(e.target.value))}>
          <option value={4}>4</option>
          <option value={8}>8</option>
          <option value={16}>16</option>
          <option value={32}>32</option>
          <option value={64}>64</option>
        </select>
        <div>
          <DeleteIcon height={16} width={16} />
        </div>
      </div>
    </div>
  );
};

export default Settings;
