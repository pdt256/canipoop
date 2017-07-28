//
//  ViewController.swift
//  canipoop
//
//  Created by Raghav Mangrola on 7/27/17.
//  Copyright © 2017 Raghav Mangrola. All rights reserved.
//

import UIKit
import FirebaseDatabase

class CanIPoopViewController: UIViewController {

    private let closedString = "👎"
    private let openString = "👍"

    @IBOutlet weak var upstairsView: UIView!
    @IBOutlet weak var upstairsLabel: UILabel!
    @IBOutlet weak var downstairsView: UIView!
    @IBOutlet weak var downstairsLabel: UILabel!

    private var ref: DatabaseReference!

    override var prefersStatusBarHidden: Bool {
        return true
    }

    override func viewDidLoad() {
        super.viewDidLoad()

        ref = Database.database().reference()
        observeDataFromFirebase()
    }

    private func observeDataFromFirebase() {
        ref.observe(DataEventType.value, with: { (snapshot) in
            let postDict = snapshot.value as? [String: AnyObject] ?? [:]

            guard let officeOne = postDict["office1"] as? [String: Any],
                let downstairsBathroom = officeOne["br1"] as? [String: Any],
                let upstairsBathroom = officeOne["br2"] as? [String: Any],
                let isDownstairsOpen = downstairsBathroom["isOpen"] as? NSNumber,
                let isUpstairsOpen = upstairsBathroom["isOpen"] as? NSNumber else {
                    print("oh no")
                    return
            }

            self.updateBathroomStatus(isDownstairsOpen.boolValue, isUpstairsOpen.boolValue)
        })
    }

    private func updateBathroomStatus(_ isDownstairsBathroomOpen: Bool, _ isUpstairsBathroomOpen: Bool) {
        if isDownstairsBathroomOpen {
            downstairsView.backgroundColor = UIColor(red:0.29, green:0.91, blue:0.49, alpha:1.00)
            downstairsLabel.text = openString
        } else {
            downstairsView.backgroundColor = UIColor(red:0.22, green:0.19, blue:0.39, alpha:1.00)
            downstairsLabel.text = closedString
        }

        if isUpstairsBathroomOpen {
            upstairsView.backgroundColor = UIColor(red:0.29, green:0.91, blue:0.49, alpha:1.00)
            upstairsLabel.text = openString
        } else {
            upstairsView.backgroundColor = UIColor(red:0.22, green:0.19, blue:0.39, alpha:1.00)
            upstairsLabel.text = closedString
        }
    }
}
