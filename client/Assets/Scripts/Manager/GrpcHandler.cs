using System.Collections;
using System.Collections.Generic;
using NetLib;
using UnityEngine;

public class GrpcHandler : MonoBehaviour
{
    // Start is called before the first frame update

    private void OnDestroy()
    {
//        UnityEngine.UI.InputField 
        RPCManager.Instance.Dispose();
    }
}
