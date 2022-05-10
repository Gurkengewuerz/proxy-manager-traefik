import {getAudit} from "../../service/Audit";
import {useEffect, useState} from "react";
import moment from 'moment';

const Comp = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    getAudit().then(resp => setData(resp.data));
  }, []);

  return <div className="row justify-content-center">
    <div className="col-8">
      <div className="card">
        <div className="card-body">
          <div className="divide-y">
            {data.map(value => (<div>
              <div className="row">
                <div className="col-auto">
                  <span className="avatar">{value.user || "UN"}</span>
                </div>
                <div className="col">
                  <div className="text-truncate">
                    {value.translated}
                  </div>
                  <div className="text-muted" title={moment(value.CreatedAt).format("LLLL")}>{moment(value.CreatedAt).fromNow()}</div>
                </div>
              </div>
            </div>))}
          </div>
        </div>
      </div>
    </div>
  </div>
}
export const title = "Audit";
export default Comp;